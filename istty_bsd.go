//go:build !purego && (darwin || freebsd || openbsd || netbsd || dragonfly)

package istty

import (
	"syscall"
	"unsafe"
)

func isTerminal(fd uintptr) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL,
		fd, syscall.TIOCGETA,
		uintptr(unsafe.Pointer(&termios)),
		0, 0, 0)
	return err == 0
}
