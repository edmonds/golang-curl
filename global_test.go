package curl

import "testing"

func TestGlobal(t *testing.T) {
    GlobalCleanup()

    if err := GlobalInit(GLOBAL_NOTHING); err != nil {
        t.Error("GlobalInit() with GLOBAL_NOTHING failed")
    }
    GlobalCleanup()

    if err := GlobalInit(GLOBAL_ALL | GLOBAL_ACK_EINTR); err != nil {
        t.Error("GlobalInit() with GLOBAL_ALL | GLOBAL_ACK_EINTR failed")
    }
    GlobalCleanup()
}
