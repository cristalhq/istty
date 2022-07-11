package istty

// IsTerminal returns true if the given file descriptor is terminal.
func IsTerminal(fd uintptr) bool {
	return isTerminal(fd)
}
