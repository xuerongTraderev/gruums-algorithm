package algorithm

import (
	"fmt"
	"testing"
)

func TestFloodFill(t *testing.T) {
	m := [][]int{
		{1,1,1,1},
		{1,0,0,1},
		{1,0,0,1},
		{1,1,1,1},
	}
	FloodFill(m, &D2ArrayIndex{0,3}, 3)
	fmt.Println("end")
	if 1 == 0 {
		t.Errorf("FirstRecurringCharacter fail")
	}
}