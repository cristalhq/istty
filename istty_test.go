package istty

import (
	"os"
	"testing"
)

func TestTTY(t *testing.T) {
	isTTY := IsTerminal(os.Stdin.Fd())
	t.Logf("isTTY: %v", isTTY)
}
