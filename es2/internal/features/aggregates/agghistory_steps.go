package aggregates

import . "github.com/lsegal/gucumber"

func init() {
	Given(`^an event sourced aggregate's event history$`, func() {

	})

	When(`^I instantiate the aggregate from its history$`, func() {

	})

	Then(`^the instance state is correct$`, func() {

	})

	And(`^there are no uncommitted events$`, func() {
		T.Skip() // pending
	})

}
