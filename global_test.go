package curl

import "testing"

func TestGlobal(t *testing.T) {
    var err *CurlError

    GlobalCleanup()

    globalFlags = GLOBAL_NOTHING
    if err = GlobalInit(); err != nil {
        t.Error("GlobalInit() with GLOBAL_NOTHING failed")
    }
    GlobalCleanup()

    globalFlags = GLOBAL_ALL | GLOBAL_ACK_EINTR
    if err = GlobalInit(); err != nil {
        t.Error("GlobalInit() with GLOBAL_ALL | GLOBAL_ACK_EINTR failed")
    }
    GlobalCleanup()
}
