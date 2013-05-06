package curl

import "io"

func IOWriterCallback(easy *Easy, buf []byte, userData interface{}) int {
    w := userData.(io.Writer)
    if _, err := w.Write(buf); err != nil {
        return 0
    }
    return len(buf)
}
