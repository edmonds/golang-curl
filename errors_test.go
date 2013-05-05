package curl

import "testing"

func TestErrors(t *testing.T) {
    want_string := "Socket not ready for send/recv"
    got_string := NewError(E_AGAIN).Error()
    if want_string != got_string {
        t.Errorf("want_string (%q) != got_string (%q)\n", want_string, got_string)
    }
    if NewError(E_OK) != nil {
        t.Error("E_OK is not nil")
    }
    if NewError(E_FAILED_INIT) == nil {
        t.Error("E_FAILED_INIT is nil")
    }
}
