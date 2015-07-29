package primefactors

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePrimes(t *testing.T) {

	pf := PrimeFactors(918125)
	assert.Equal(t, 6,len(pf))
	assert.Equal(t, 5, pf[0])
	assert.Equal(t, 5, pf[1])
	assert.Equal(t, 5, pf[2])
	assert.Equal(t, 5, pf[3])
	assert.Equal(t, 13, pf[4])
	assert.Equal(t, 113, pf[5])
	fmt.Println(pf)

	pf = PrimeFactors(8125)
	assert.Equal(t, 5,len(pf))
	assert.Equal(t, 5, pf[0])
	assert.Equal(t, 5, pf[1])
	assert.Equal(t, 5, pf[2])
	assert.Equal(t, 5, pf[3])
	assert.Equal(t, 13, pf[4])
	fmt.Println(pf)

	pf = PrimeFactors(315)
	assert.Equal(t, 4,len(pf))
	assert.Equal(t, 3, pf[0])
	assert.Equal(t, 3, pf[1])
	assert.Equal(t, 5, pf[2])
	assert.Equal(t, 7, pf[3])
	fmt.Println(pf)
	
}