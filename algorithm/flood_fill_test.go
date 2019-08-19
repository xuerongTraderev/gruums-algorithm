package algorithm

import (
	"testing"
)

func compare2DArray(a [][]int, b [][]int) bool {
	for r := 0 ; r < len(a) ; r++ {
		for j := 0 ; j < len(a[r]) ; j++ {
			if a[r][j] != b[r][j] {
				return false
			}
		}
	}
	return true;
}

func TestFloodFill(t *testing.T) {
	m := [][]int{
		{1,1,1,1},
		{1,0,0,1},
		{1,0,0,1},
		{1,1,1,1},
	}
	FloodFill(m, &D2ArrayIndex{0,3}, 3)
	m2 := [][]int{
		{3,3,3,3},
		{3,0,0,3},
		{3,0,0,3},
		{3,3,3,3},
	}
	if !compare2DArray(m, m2) {
		t.Errorf("FirstRecurringCharacter fail")
	}
}