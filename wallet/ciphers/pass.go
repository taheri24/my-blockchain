package ciphers

import "io"

type repeatedStringReader []byte

func NewReader(pass string) io.Reader {
	return repeatedStringReader([]byte(pass))
}
func (r repeatedStringReader) Read(p []byte) (n int, err error) {
	c, count := len(r), len(p)
	for i := 0; i < count; i++ {
		p[i] = r[i%c]
	}

	return count, nil
}
