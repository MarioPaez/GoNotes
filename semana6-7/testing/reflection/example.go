package reflection

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

func try() {
	var r io.Reader
	r = os.Stdin
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
	// Y cualquier que implemente Read se puede igualar a r
}
