package curl

import "fmt"

//#cgo pkg-config: libcurl
//#include <curl/curl.h>
import "C"

/* Constants */

const (
    GLOBAL_ALL = C.CURL_GLOBAL_ALL
    GLOBAL_SSL = C.CURL_GLOBAL_SSL
    GLOBAL_WIN32 = C.CURL_GLOBAL_WIN32
    GLOBAL_NOTHING = C.CURL_GLOBAL_NOTHING
    GLOBAL_DEFAULT = C.CURL_GLOBAL_DEFAULT
    GLOBAL_ACK_EINTR = C.CURL_GLOBAL_ACK_EINTR
)
type GlobalFlags int

/* Variables */

var globalFlags GlobalFlags = GLOBAL_DEFAULT

/* Functions */

func GlobalInit() *CurlError {
    return NewError(C.curl_global_init(C.long(globalFlags)))
}

func GlobalCleanup() {
    C.curl_global_cleanup()
}

func init() {
    if err := GlobalInit(); err != nil {
        panic(fmt.Sprintf("unable to initialize libcurl: %s", err))
    }
}
