package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (s spaceEraser) Read(a []byte) (int, error) {
	e, err := s.r.Read(a)
	//Remove all spaces by replacing them by nothing
	b := bytes.ReplaceAll(a, []byte(" "), []byte(""))
	//Some withe spaces can still be remaining
	b = bytes.Trim(b, "\x00")
	//Need to update the length of the byte
	e, err = s.r.Read(b)
	e = copy(a, b)
	return e, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"
