package timer

import (
	"testing"
)

func TestPostitiveDuration(t *testing.T) {
	at := NewAPITimer("foo")
	at.Stop(nil)
	if at.Time == 0 {
		t.Fail()
	}
}