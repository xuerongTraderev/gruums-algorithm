package algorithm

import "testing"

func Test(t *testing.T) {
	if FirstRecurringCharacter() != "aaaa" {
		t.Errorf("want %q", "aaaa")
	}
}