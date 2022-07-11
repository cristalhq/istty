//go:build purego

package istty

func isTerminal(fd uintptr) bool {
	return false
}
