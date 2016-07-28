package timer

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
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

func TestMultiBackendRecordings(t *testing.T) {
	at := NewAPITimer("foo")
	c1 := at.StartContributor("c1")
	c2 := at.StartContributor("c2")
	c3 := at.StartContributor("c3")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		be1 := c3.StartServiceCall("workflo")
		be1.End()
	}()

	go func() {
		defer wg.Done()
		be2 := c3.StartServiceCall("doc munger")
		be2.End()
	}()

	wg.Wait()

	c3.End(nil)

	c2.End(nil)
	c1.End(nil)
	at.Stop(nil)

	if at.Error != nil {
		t.Fail()
	}

	if c1.Time <= 0 || c2.Time <= 0 || c3.Time <= 0 {
		t.Fail()
	}

	if at.ErrorFree == false {
		t.Fail()
	}

	if len(c3.ServiceCalls) != 2 {
		t.Fail()
	}

	println(at.ToJSONString())
}

func MultiGoRoutines(t *testing.T) {
	for i := 1; i < 1000; i++ {
		at := NewAPITimer("foo")
		c1 := at.StartContributor("c1")

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			c1 := at.StartContributor("c1")

			go func() {
				defer wg.Done()
				c2 := at.StartContributor("c2")
				time.Sleep(50 * time.Millisecond)
				c2.End(nil)
			}()

			time.Sleep(15 * time.Millisecond)
			c1.End(nil)

		}()

		wg.Wait()

		c1.End(nil)

		at.Stop(nil)

		wg.Add(1)

		go func() {
			fmt.Println(at.ToJSONString())
			wg.Done()
		}()

		wg.Wait()
	}
}
