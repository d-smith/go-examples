package aggregates

import (
    . "github.com/lsegal/gucumber"
	"github.com/stretchr/testify/assert"
	"github.com/d-smith/go-examples/es2/sample"
)



var u1, u2 *sample.User

func init() {
	Given(`^an event sourced aggregate$`, func() {		
	})

	When(`^I create new instances$`, func() {
		u1 = sample.NewUser()
		u2 = sample.NewUser()
	})

	Then(`^the instances have unique IDs$`, func() {
		assert.NotEqual(T,u1.ID, u2.ID,"IDs were not unique")
	})

}