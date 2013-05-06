package curl

/*
#include <stdlib.h>
#include <curl/curl.h>
#include "cwrap.h"

static CURLcode
my_setoptc(CURL *handle, CURLoption option, char *param) {
    return curl_easy_setopt(handle, option, param);
}

static CURLcode
my_setoptl(CURL *handle, CURLoption option, long param) {
    return curl_easy_setopt(handle, option, param);
}
*/
import "C"

import "errors"
import "io"
import "fmt"
import "runtime"
import "unsafe"

type WriteFunction func(easy *Easy, buf []byte, userData interface{}) (int)

type EasyError struct {
    Code CURLcode
    String string
    EasyString string
}

type Easy struct {
    cptr unsafe.Pointer
    cerrbuf *C.char
    writeFunc WriteFunction
    writeData interface{}
}

func (err *EasyError) Error() string {
    if err.String != "" {
        return err.String
    } else {
        return err.EasyString
    }
}

func EasyInit() (easy *Easy, err *CurlError) {
    easy = new(Easy)
    if easy.cptr = C.curl_easy_init(); easy.cptr == nil {
        return nil, NewError(E_FAILED_INIT)
    }
    easy.cerrbuf = (*C.char)(C.malloc(C.CURL_ERROR_SIZE))
    err = NewError(C.my_setoptc(easy.cptr, C.CURLOPT_ERRORBUFFER, easy.cerrbuf))
    if err != nil {
        C.curl_easy_cleanup(easy.cptr)
        return nil, err
    }
    runtime.SetFinalizer(easy, func(e *Easy) { e.Cleanup() })
    return
}

func (easy *Easy) Cleanup() {
    C.curl_easy_cleanup(easy.cptr)
    C.free(unsafe.Pointer(easy.cerrbuf))
}

func (easy *Easy) Reset() {
    C.curl_easy_reset(easy.cptr)
    easy.writeFunc = nil
    easy.writeData = nil
}

func (easy *Easy) getError(code C.CURLcode) error {
    if CURLcode(code) == E_OK {
        return nil
    }
    err := new(EasyError)
    err.Code = CURLcode(code)
    err.String = C.GoString(C.curl_easy_strerror(code))
    err.EasyString = C.GoString(easy.cerrbuf)
    return err
}

func (easy *Easy) Perform() error {
    return easy.getError(C.curl_easy_perform(easy.cptr))
}

func (easy *Easy) SetOpt(opt EasyOption, param interface{}) error {
    switch {
    case opt >= C.CURLOPTTYPE_OFF_T:
        return easy.getError(C.my_setoptl(easy.cptr, C.CURLoption(opt), C.long(param.(int))))
    case opt >= C.CURLOPTTYPE_FUNCTIONPOINT:
        return errors.New("function options not supported by SetOpt()")
    case opt >= C.CURLOPTTYPE_OBJECTPOINT:
        if isStringOption(opt) {
            c_param := C.CString(param.(string))
            defer C.free(unsafe.Pointer(c_param))
            return easy.getError(C.my_setoptc(easy.cptr, C.CURLoption(opt), c_param))
        } else {
            return fmt.Errorf("option %d not supported", opt)
        }
    case opt >= C.CURLOPTTYPE_LONG:
        var val C.long
        switch t := param.(type) {
        case bool:
            if t {
                val = C.long(1)
            }
        case int:
            val = C.long(t)
        case int32:
            val = C.long(t)
        case int64:
            val = C.long(t)
        default:
            panic("CURLOPTTYPE_LONG param must be of golang type bool, int, int32, or int64")
        }
        return easy.getError(C.my_setoptl(easy.cptr, C.CURLoption(opt), val))
    }
    return nil
}

func (easy *Easy) SetOptWrite(writeFunc WriteFunction, writeData interface{}) error {
    if writeFunc == nil && writeData == nil {
        // request to disable the write callback entirely
        if err := easy.getError(C.my_set_write(nil, nil)); err != nil {
            return err
        }
        return nil
    }

    if writeFunc == nil && writeData != nil {
        // this is analogous to libcurl's fwrite() default where WRITEDATA is
        // treated as a writeable FILE * when WRITEFUNCTION is unspecified.
        // translate this to mean a golang io.Writer for which we will provide
        // an analogous default callback, and assert that the caller passed
        // such an object.
        writeFunc = IOWriterCallback
        writeData = writeData.(io.Writer)
    }

    // install the cwrap trampoline
    if err := easy.getError(C.my_set_write(easy.cptr, unsafe.Pointer(easy))); err != nil {
        return err
    }

    // save the parameters to be used by the trampoline
    easy.writeFunc = writeFunc
    easy.writeData = writeData
    return nil
}
