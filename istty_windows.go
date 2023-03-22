package istty

import (
	"syscall"
	"unsafe"
)

var (
	modkernel32        = syscall.MustLoadDLL("kernel32.dll")
	procGetConsoleMode = modkernel32.MustFindProc("GetConsoleMode")
)

func isTerminal(fd uintptr) bool {
	var st uint32
	return getConsoleMode(syscall.Handle(fd), &st) == nil
}

func getConsoleMode(hConsoleHandle syscall.Handle, lpMode *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetConsoleMode.Addr(), 2, uintptr(hConsoleHandle), uintptr(unsafe.Pointer(lpMode)), 0)
	if int(r1) == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
