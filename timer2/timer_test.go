package timer2

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const zeroDuration = 0 * time.Millisecond

func TestPostitiveDuration(t *testing.T) {
	at := NewEndToEndTimer("foo")
	time.Sleep(50 * time.Millisecond)
	at.Stop(nil)
	assert.True(t, at.EndToEndDuration() > 50*time.Millisecond)
	assert.NotEqual(t, "", at.TxnId, "Expected non-empty txn id")

	at.Kill()
}

func TestStopError(t *testing.T) {
	at := NewEndToEndTimer("foo")
	defer at.Kill()
	at.Stop(errors.New("problem"))
	assert.Equal(t, "problem", at.GetError())
	assert.False(t, at.IsErrorFree(), "Timer should indicate an error was set")

}

func TestContributors(t *testing.T) {
	at := NewEndToEndTimer("foo")
	defer at.Kill()
	c1 := at.StartContributor("c1")
	c2 := at.StartContributor("c2")
	c2.End(nil)
	c1.End(nil)
	at.Stop(nil)

	assert.Equal(t, "", at.GetError())

	assert.False(t, c1.Time() <= zeroDuration || c2.Time() <= zeroDuration)
	assert.True(t, at.IsErrorFree(), "Timer should be error free")

	fmt.Println(at.ToJSONString())

}

func TestIfContributorErrorsThenTimerErrors(t *testing.T) {
	at := NewEndToEndTimer("foo")
	defer at.Kill()
	c1 := at.StartContributor("c1")
	c1.End(errors.New("kaboom"))
	at.Stop(nil)

	assert.False(t, at.IsErrorFree(), "Expected contributor error to make timer non-error free")
	assert.Equal(t, "", at.GetError(), "No error message on top level timer expected")
	assert.Equal(t, "kaboom", c1.GetError(), "Expected kaboom as contributor error message")
}

func TestServiceCallErrorDetection(t *testing.T) {
	at := NewEndToEndTimer("foo")
	defer at.Kill()
	c1 := at.StartContributor("c1")
	sc := c1.StartServiceCall("larry", "/dev/null")
	sc.End(errors.New("blammo"))
	c1.End(nil)
	at.Stop(nil)

	assert.False(t, at.IsErrorFree(), "Expected contributor error to make timer non-error free")

	fmt.Println(at.ToJSONString())
}

//TODO - refactor error to string
