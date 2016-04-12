package timer2

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const zeroDuration = 0 * time.Millisecond

func TestPostitiveDuration(t *testing.T) {
	at := NewEndToEndTimer("foo")
	time.Sleep(50 * time.Millisecond)
	at.Stop(nil)
	if at.Duration() < 50*time.Millisecond {
		t.Fail()
	}

	at.Kill()
}

func TestStopError(t *testing.T) {
	at := NewEndToEndTimer("foo")
	at.Stop(errors.New("problem"))
	assert.Equal(t, "problem", at.Error())
}

func TestContributors(t *testing.T) {
	at := NewEndToEndTimer("foo")
	c1 := at.StartContributor("c1")
	c2 := at.StartContributor("c2")
	c2.End(nil)
	c1.End(nil)
	at.Stop(nil)

	assert.Equal(t, "", at.Error())

	if c1.Time() <= zeroDuration || c2.Time() <= zeroDuration {
		t.Fail()
	}

	if at.ErrorFree() == false {
		t.Fail()
	}

}

func TestIfContributorErrorsThenTimerErrors(t *testing.T) {

}

func TestMultiBackendRecordings(t *testing.T) {

}
