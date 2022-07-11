//go:build !purego && linux

package istty

import (
	"syscall"
	"unsafe"
)

func isTerminal(fd uintptr) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL,
		fd, syscall.TCGETS,
		uintptr(unsafe.Pointer(&termios)),
		0, 0, 0)
	return err == 0
}
