package aggregates

import (
	"github.com/d-smith/go-examples/es2/sample"
	. "github.com/lsegal/gucumber"
	"github.com/stretchr/testify/assert"
)

var u1, u2 *sample.User

func init() {
	Given(`^an event sourced aggregate$`, func() {
	})

	When(`^I create new instances$`, func() {
		u1, _ = sample.NewUser("first", "last", "one@foo.com")
		u2, _ = sample.NewUser("first", "last", "two@foo.com")

		assert.NotNil(T, u1)
		assert.NotNil(T, u2)
	})

	Then(`^the instances have unique IDs$`, func() {
		assert.NotEqual(T, u1.ID, u2.ID, "IDs were not unique")
	})

	And(`^there's an uncommitted event$`, func() {
		assert.Equal(T, 1, len(u1.Aggregate.Events))
	})
}
