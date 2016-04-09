package timer2

import (
	"testing"
	"time"
)

func TestPostitiveDuration(t *testing.T) {
	at := NewEndToEndTimer("foo")
	time.Sleep(50 * time.Millisecond)
	at.Stop(nil)
	if at.Duration() < 50*time.Millisecond {
		t.Fail()
	}

	at.Kill()
}

func TestContributors(t *testing.T) {

}

func TestIfContributorErrorsThenTimerErrors(t *testing.T) {

}

func TestMultiBackendRecordings(t *testing.T) {

}
