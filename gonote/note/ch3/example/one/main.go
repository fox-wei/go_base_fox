package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

type capitalzedReader struct {
	r io.Reader
}

func (r *capitalzedReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)

	if err != nil {
		return 0, err
	}

	q := bytes.ToUpper(p) //?大写
	copy(p, q)
	return n, err
}

//*装饰函数
func CapReader(r io.Reader) io.Reader {
	return &capitalzedReader{r: r}
}

func main() {
	r := strings.NewReader("Hello ying!")
	lr := CapReader(io.LimitReader(r, 5))
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}
