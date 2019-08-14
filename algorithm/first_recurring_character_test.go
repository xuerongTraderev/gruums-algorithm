package algorithm

import (
	"testing"
)
func TestFirstRecurringCharacter(t *testing.T) {
	if FirstRecurringCharacter("ad呜呜sa") != '呜' {
		t.Errorf("FirstRecurringCharacter fail")
	}
}