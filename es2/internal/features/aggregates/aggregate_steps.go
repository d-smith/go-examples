package aggregates

import (
    . "github.com/lsegal/gucumber"
	"github.com/stretchr/testify/assert"
    "github.com/d-smith/go-examples/es2"
)

type TestAgg struct {
    *es2.Aggregate
}

func NewTestAgg() *TestAgg {
	var testAgg = new(TestAgg)
	testAgg.Aggregate = es2.NewAggregate()
	return testAgg
}

var ta1, ta2 *TestAgg

func init() {
	Given(`^an event sourced aggregate$`, func() {		
	})

	When(`^I create new instances$`, func() {
		ta1 = NewTestAgg()
		ta2 = NewTestAgg()
	})

	Then(`^the instances have unique IDs$`, func() {
		assert.NotEqual(T,ta1.ID, ta2.ID,"IDs were not unique")
	})

}