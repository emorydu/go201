package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("hello, emorydu!\n")
	lr := io.LimitReader(r, 4)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}

	println()

	r = strings.NewReader("hello, emorydu!\n")
	r1 := CapReader(io.LimitReader(r, 4))
	if _, err := io.Copy(os.Stdout, r1); err != nil {
		log.Fatal(err)
	}
}

func CapReader(r io.Reader) io.Reader {
	return &capitalizedReader{r: r}
}

type capitalizedReader struct {
	r io.Reader
}

func (r *capitalizedReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return 0, err
	}

	q := bytes.ToUpper(p)
	for i, v := range q {
		p[i] = v
	}
	return n, err
}
