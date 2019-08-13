package algorithm

import (
	"testing"
)
func Test(t *testing.T) {
	if FirstRecurringCharacter("ad呜呜sa") != '呜' {
		t.Errorf("FirstRecurringCharacter fail")
	}
}