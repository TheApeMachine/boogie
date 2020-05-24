package apeterm

import (
	"io"
	"os"
	"sync"
)

// UI is an abstraction around the rendering of visual elements
// in the terminal window, as well as handling input.
//
// Some of this code was shamelessly ripped from:
// Taichi Nakashima
// https://github.com/tcnksm/go-input
type UI struct {
	Writer io.Writer
	Reader io.Reader
	once   sync.Once
}

func NewUI() *UI {
	return &UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}
}

// ReadInputByte, bit brute-force.
// TODO: Please get the cast out of here!
func (ui *UI) ReadInputByte() (string, error) {
	var rBuf []byte

	for {
		var buf [1]byte
		_, err := ui.Reader.(*os.File).Read(buf[:])

		if err != nil && err != io.EOF {
			return "", err
		}

		if buf[0] == 3 {
			return "", nil
		}

		rBuf = append(rBuf, buf[0])

		if len(rBuf) > 0 {
			return string(rBuf), nil
		}
	}

	return string(rBuf), nil
}
