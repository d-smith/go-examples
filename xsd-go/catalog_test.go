package main

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestComplexTypeCatalog(t *testing.T) {
	bs, err := ioutil.ReadFile("XtracElements.xsd")
	assert.Nil(t, err)

	s := new(Schema)

	err = xml.Unmarshal(bs, &s)
	assert.Nil(t, err)

	catalog := NewCatalog()
	catalog.CatalogComplexTypes(s)

	ct, err := catalog.LookupComplexType("LoginInfo_T")
	assert.Nil(t, err)
	assert.Equal(t, "LoginInfo_T", ct.Name)
	assert.Equal(t, 7, len(ct.Sequence))

	for _, e := range ct.Sequence {
		fmt.Printf("ref: %s\n", e.Ref)
	}

	//Next - add a predicate to see if we are referencing a simple type or a complex type.
}
