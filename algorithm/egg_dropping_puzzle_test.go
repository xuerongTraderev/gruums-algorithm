package algorithm

import (
	"testing"
)
func TestEggDroppingPuzzle(t *testing.T) {
	if EggDroppingPuzzle(2, 36) != 8 {
		t.Errorf("EggDroppingPuzzle fail")
	}
}