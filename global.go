package curl

import "fmt"

//#cgo pkg-config: libcurl
//#include <curl/curl.h>
import "C"

const (
    GLOBAL_ALL = C.CURL_GLOBAL_ALL
    GLOBAL_SSL = C.CURL_GLOBAL_SSL
    GLOBAL_WIN32 = C.CURL_GLOBAL_WIN32
    GLOBAL_NOTHING = C.CURL_GLOBAL_NOTHING
    GLOBAL_DEFAULT = C.CURL_GLOBAL_DEFAULT
    GLOBAL_ACK_EINTR = C.CURL_GLOBAL_ACK_EINTR
)

func GlobalInit(flags int) *CurlError {
    return NewError(C.curl_global_init(C.long(flags)))
}

func GlobalCleanup() {
    C.curl_global_cleanup()
}

func init() {
    if err := GlobalInit(GLOBAL_DEFAULT); err != nil {
        panic(fmt.Sprintf("libcurl failed to initialize: %s", err))
    }
}
