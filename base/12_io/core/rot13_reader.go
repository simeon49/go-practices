package core

import (
	"io"
	"os"
	"strings"
)

// rot13 代换密码 参考: https://en.wikipedia.org/wiki/ROT13

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		ch := b[i]
		if (ch >= 'A' && ch <= 'M') || (ch >= 'a' && ch <= 'm') {
			ch += 13
		} else if (ch >= 'N' && ch <= 'Z') || (ch >= 'n' && ch <= 'z') {
			ch -= 13
		}
		b[i] = ch
	}
	return n, nil
}

func Rot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
