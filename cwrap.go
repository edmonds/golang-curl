package curl

//#include <stdlib.h>
//#cgo pkg-config: libcurl
//#include <curl/curl.h>
import "C"

import "unsafe"

//export go_write_callback
func go_write_callback(ptr *C.char, size C.size_t, nmemb C.size_t, user unsafe.Pointer) (C.size_t) {
    easy := (*Easy)(user)
    if easy == nil {
        panic("easy == nil")
    }
    if easy.writeFunc == nil {
        panic("easy.write == nil")
    }
    buf := C.GoBytes(unsafe.Pointer(ptr), C.int(size * nmemb))
    ret := easy.writeFunc(easy, buf, easy.writeData)
    return C.size_t(ret)
}
