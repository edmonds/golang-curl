package curl

//#cgo pkg-config: libcurl
//#include <curl/curl.h>
import "C"

/* Constants */

const (
    E_ABORTED_BY_CALLBACK = C.CURLE_ABORTED_BY_CALLBACK
    E_AGAIN = C.CURLE_AGAIN
    E_BAD_CONTENT_ENCODING = C.CURLE_BAD_CONTENT_ENCODING
    E_BAD_DOWNLOAD_RESUME = C.CURLE_BAD_DOWNLOAD_RESUME
    E_BAD_FUNCTION_ARGUMENT = C.CURLE_BAD_FUNCTION_ARGUMENT
    E_CHUNK_FAILED = C.CURLE_CHUNK_FAILED
    E_CONV_FAILED = C.CURLE_CONV_FAILED
    E_CONV_REQD = C.CURLE_CONV_REQD
    E_COULDNT_CONNECT = C.CURLE_COULDNT_CONNECT
    E_COULDNT_RESOLVE_HOST = C.CURLE_COULDNT_RESOLVE_HOST
    E_COULDNT_RESOLVE_PROXY = C.CURLE_COULDNT_RESOLVE_PROXY
    E_FAILED_INIT = C.CURLE_FAILED_INIT
    E_FILESIZE_EXCEEDED = C.CURLE_FILESIZE_EXCEEDED
    E_FILE_COULDNT_READ_FILE = C.CURLE_FILE_COULDNT_READ_FILE
    E_FTP_ACCEPT_FAILED = C.CURLE_FTP_ACCEPT_FAILED
    E_FTP_ACCEPT_TIMEOUT = C.CURLE_FTP_ACCEPT_TIMEOUT
    E_FTP_BAD_FILE_LIST = C.CURLE_FTP_BAD_FILE_LIST
    E_FTP_CANT_GET_HOST = C.CURLE_FTP_CANT_GET_HOST
    E_FTP_COULDNT_RETR_FILE = C.CURLE_FTP_COULDNT_RETR_FILE
    E_FTP_COULDNT_SET_TYPE = C.CURLE_FTP_COULDNT_SET_TYPE
    E_FTP_COULDNT_USE_REST = C.CURLE_FTP_COULDNT_USE_REST
    E_FTP_PORT_FAILED = C.CURLE_FTP_PORT_FAILED
    E_FTP_PRET_FAILED = C.CURLE_FTP_PRET_FAILED
    E_FTP_WEIRD_227_FORMAT = C.CURLE_FTP_WEIRD_227_FORMAT
    E_FTP_WEIRD_PASS_REPLY = C.CURLE_FTP_WEIRD_PASS_REPLY
    E_FTP_WEIRD_PASV_REPLY = C.CURLE_FTP_WEIRD_PASV_REPLY
    E_FTP_WEIRD_SERVER_REPLY = C.CURLE_FTP_WEIRD_SERVER_REPLY
    E_FUNCTION_NOT_FOUND = C.CURLE_FUNCTION_NOT_FOUND
    E_GOT_NOTHING = C.CURLE_GOT_NOTHING
    E_HTTP_POST_ERROR = C.CURLE_HTTP_POST_ERROR
    E_HTTP_RETURNED_ERROR = C.CURLE_HTTP_RETURNED_ERROR
    E_INTERFACE_FAILED = C.CURLE_INTERFACE_FAILED
    E_LDAP_CANNOT_BIND = C.CURLE_LDAP_CANNOT_BIND
    E_LDAP_INVALID_URL = C.CURLE_LDAP_INVALID_URL
    E_LDAP_SEARCH_FAILED = C.CURLE_LDAP_SEARCH_FAILED
    E_LOGIN_DENIED = C.CURLE_LOGIN_DENIED
    E_NOT_BUILT_IN = C.CURLE_NOT_BUILT_IN
    E_NO_CONNECTION_AVAILABLE = C.CURLE_NO_CONNECTION_AVAILABLE
    E_OK = C.CURLE_OK
    E_OPERATION_TIMEDOUT = C.CURLE_OPERATION_TIMEDOUT
    E_OUT_OF_MEMORY = C.CURLE_OUT_OF_MEMORY
    E_PARTIAL_FILE = C.CURLE_PARTIAL_FILE
    E_PEER_FAILED_VERIFICATION = C.CURLE_PEER_FAILED_VERIFICATION
    E_QUOTE_ERROR = C.CURLE_QUOTE_ERROR
    E_RANGE_ERROR = C.CURLE_RANGE_ERROR
    E_READ_ERROR = C.CURLE_READ_ERROR
    E_RECV_ERROR = C.CURLE_RECV_ERROR
    E_REMOTE_ACCESS_DENIED = C.CURLE_REMOTE_ACCESS_DENIED
    E_REMOTE_DISK_FULL = C.CURLE_REMOTE_DISK_FULL
    E_REMOTE_FILE_EXISTS = C.CURLE_REMOTE_FILE_EXISTS
    E_REMOTE_FILE_NOT_FOUND = C.CURLE_REMOTE_FILE_NOT_FOUND
    E_RTSP_CSEQ_ERROR = C.CURLE_RTSP_CSEQ_ERROR
    E_RTSP_SESSION_ERROR = C.CURLE_RTSP_SESSION_ERROR
    E_SEND_ERROR = C.CURLE_SEND_ERROR
    E_SEND_FAIL_REWIND = C.CURLE_SEND_FAIL_REWIND
    E_SSH = C.CURLE_SSH
    E_SSL_CACERT = C.CURLE_SSL_CACERT
    E_SSL_CACERT_BADFILE = C.CURLE_SSL_CACERT_BADFILE
    E_SSL_CERTPROBLEM = C.CURLE_SSL_CERTPROBLEM
    E_SSL_CIPHER = C.CURLE_SSL_CIPHER
    E_SSL_CONNECT_ERROR = C.CURLE_SSL_CONNECT_ERROR
    E_SSL_CRL_BADFILE = C.CURLE_SSL_CRL_BADFILE
    E_SSL_ENGINE_INITFAILED = C.CURLE_SSL_ENGINE_INITFAILED
    E_SSL_ENGINE_NOTFOUND = C.CURLE_SSL_ENGINE_NOTFOUND
    E_SSL_ENGINE_SETFAILED = C.CURLE_SSL_ENGINE_SETFAILED
    E_SSL_ISSUER_ERROR = C.CURLE_SSL_ISSUER_ERROR
    E_SSL_SHUTDOWN_FAILED = C.CURLE_SSL_SHUTDOWN_FAILED
    E_TELNET_OPTION_SYNTAX = C.CURLE_TELNET_OPTION_SYNTAX
    E_TFTP_ILLEGAL = C.CURLE_TFTP_ILLEGAL
    E_TFTP_NOSUCHUSER = C.CURLE_TFTP_NOSUCHUSER
    E_TFTP_NOTFOUND = C.CURLE_TFTP_NOTFOUND
    E_TFTP_PERM = C.CURLE_TFTP_PERM
    E_TFTP_UNKNOWNID = C.CURLE_TFTP_UNKNOWNID
    E_TOO_MANY_REDIRECTS = C.CURLE_TOO_MANY_REDIRECTS
    E_UNKNOWN_OPTION = C.CURLE_UNKNOWN_OPTION
    E_UNSUPPORTED_PROTOCOL = C.CURLE_UNSUPPORTED_PROTOCOL
    E_UPLOAD_FAILED = C.CURLE_UPLOAD_FAILED
    E_URL_MALFORMAT = C.CURLE_URL_MALFORMAT
    E_USE_SSL_FAILED = C.CURLE_USE_SSL_FAILED
    E_WRITE_ERROR = C.CURLE_WRITE_ERROR
)

const (
    M_BAD_EASY_HANDLE = C.CURLM_BAD_EASY_HANDLE
    M_BAD_HANDLE = C.CURLM_BAD_HANDLE
    M_BAD_SOCKET = C.CURLM_BAD_SOCKET
    M_CALL_MULTI_PERFORM = C.CURLM_CALL_MULTI_PERFORM
    M_INTERNAL_ERROR = C.CURLM_INTERNAL_ERROR
    M_OK = C.CURLM_OK
    M_OUT_OF_MEMORY = C.CURLM_OUT_OF_MEMORY
    M_UNKNOWN_OPTION = C.CURLM_UNKNOWN_OPTION
)

const (
    SHE_BAD_OPTION = C.CURLSHE_BAD_OPTION
    SHE_INVALID = C.CURLSHE_INVALID
    SHE_IN_USE = C.CURLSHE_IN_USE
    SHE_NOMEM = C.CURLSHE_NOMEM
    SHE_NOT_BUILT_IN = C.CURLSHE_NOT_BUILT_IN
    SHE_OK = C.CURLSHE_OK
)

/* Types */

type CURLcode int
type CURLMcode int
type CURLSHcode int

type CurlError struct {
    Code CURLcode
    String string
}

type MultiError struct {
    Code CURLMcode
    String string
}

type ShareError struct {
    Code CURLSHcode
    String string
}

/* Functions */

func NewError(code C.CURLcode) (err *CurlError) {
    if CURLcode(code) == E_OK {
        return nil
    }
    err = new(CurlError)
    err.Code = CURLcode(code)
    err.String = C.GoString(C.curl_easy_strerror(C.CURLcode(err.Code)))
    return err
}

func (err *CurlError) Error() string {
    return err.String
}
