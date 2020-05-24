package apeterm

import (
	"syscall"
	"unsafe"
)

type Window struct {
	Row uint16
	Col uint16
	X   uint16
	Y   uint16
}

func NewWindow() *Window {
	return &Window{}
}

func (win *Window) getWinSize(fd int) {
	retCode, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL, uintptr(fd),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(win)),
	)

	if int(retCode) == -1 {
		panic(errno)
	}
}
