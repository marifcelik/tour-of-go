package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/tour/reader"
)

type MyReader struct{}
type MyReaderErr float64

func (e MyReaderErr) Error() string {
	return fmt.Sprintf("max capacity error, cap: %v", float64(e))
}

func (r MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return 0, MyReaderErr(cap(b))
	}

	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	// n, err := r.r.Read(b)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {

		}
	}

	return 1, nil
}

func readers() {
	fmt.Println("\n-- readers --")
	r := strings.NewReader("hello reader !")

	b := make([]byte, 6)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	// reader exercise
	fmt.Println("\nreader exercise")
	reader.Validate(MyReader{})

	// TODO: implement rot13Reader
	// s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// rotR := rot13Reader(s)
	// io.Copy(os.Stdout, &rotR)
}
