package core

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/tour/reader"
)

// ------------------------- Reader01 -------------------------

func Reader01() {
	r := strings.NewReader("This is a test, you can read.")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("%v (%v bytes) err: %v\n", b, n, err)
		fmt.Printf("b[:n] %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// ------------------------- Reader02 -------------------------

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	length := len(b)
	for i := 0; i < length; i++ {
		b[i] = 'A'
	}
	return length, nil
}

func Reader02() {
	reader.Validate(MyReader{})
}
