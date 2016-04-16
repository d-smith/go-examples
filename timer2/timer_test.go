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
	assert.True(t, at.Duration() > 50*time.Millisecond)
	assert.NotEqual(t, "", at.txnId, "Expected non-empty txn id")

	at.Kill()
}

func TestStopError(t *testing.T) {
	at := NewEndToEndTimer("foo")
	at.Stop(errors.New("problem"))
	assert.Equal(t, "problem", at.Error())
	assert.False(t, at.ErrorFree(), "Timer should indicate an error was set")
}

func TestContributors(t *testing.T) {
	at := NewEndToEndTimer("foo")
	c1 := at.StartContributor("c1")
	c2 := at.StartContributor("c2")
	c2.End(nil)
	c1.End(nil)
	at.Stop(nil)

	assert.Equal(t, "", at.Error())

	assert.False(t, c1.Time() <= zeroDuration || c2.Time() <= zeroDuration)
	assert.True(t, at.ErrorFree(), "Timer should be error free")

	at.Kill()

}

func TestIfContributorErrorsThenTimerErrors(t *testing.T) {
	at := NewEndToEndTimer("foo")
	c1 := at.StartContributor("c1")
	c1.End(errors.New("kaboom"))
	at.Stop(nil)

	assert.False(t, at.ErrorFree(), "Expected contributor error to make timer non-error free")
	assert.Equal(t, "", at.Error(), "No error message on top level timer expected")
	assert.Equal(t, "kaboom", c1.Error(), "Expected kaboom as contributor error message")
}

//Next - service start and stop, toJSON
