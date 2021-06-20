package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	reader io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.reader.Read(b)
	if err == nil {
		for idx, char := range b {
			switch {
			case char >= 'A' && char <= 'Z':
				char += 13
				if char > 'Z' {
					char -= 26
				}
				b[idx] = char
			case char >= 'a' && char <= 'z':
				char += 13
				if char > 'z' {
					char -= 26
				}
				b[idx] = char
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

