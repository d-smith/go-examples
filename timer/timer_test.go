package timer

import (
	"errors"
	"testing"
)

func TestPostitiveDuration(t *testing.T) {
	at := NewAPITimer("foo")
	at.Stop(nil)
	if at.Time == 0 {
		t.Fail()
	}

}

func TestContributors(t *testing.T) {
	at := NewAPITimer("foo")
	c1 := at.StartContributor("c1")
	c2 := at.StartContributor("c2")
	c2.End(nil)
	c1.End(nil)
	at.Stop(nil)

	if at.Error != nil {
		t.Fail()
	}

	if c1.Time <= 0 || c2.Time <= 0 {
		t.Fail()
	}

	if at.ErrorFree == false {
		t.Fail()
	}
}

func TestIfContributorErrorsThenTimerErrors(t *testing.T) {
	at := NewAPITimer("foo")
	c1 := at.StartContributor("c1")
	c2 := at.StartContributor("c2")
	c2.End(errors.New("oh whoops"))
	c1.End(nil)
	at.Stop(nil)

	if at.Error != nil {
		t.Fail()
	}

	if len(at.ContributorErrors()) != 1 {
		t.Fail()
	}

	if at.ErrorFree == true {
		t.Fail()
	}
}
