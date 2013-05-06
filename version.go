package curl

//#cgo pkg-config: libcurl
//#include <curl/curl.h>
import "C"

// curl_version()
func Version() string {
    return C.GoString(C.curl_version())
}
