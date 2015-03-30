package main

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func loadSchemaDefs(t *testing.T, xsdFileName string) *Schema {
	bs, err := ioutil.ReadFile(xsdFileName)
	assert.Nil(t, err)

	s := new(Schema)
	err = xml.Unmarshal(bs, &s)
	assert.Nil(t, err)
	return s
}

func TestComplexTypeCatalog(t *testing.T) {
	elementsSchema := loadSchemaDefs(t, "XtracElements.xsd")
	typesSchema := loadSchemaDefs(t, "XtracTypes.xsd")

	catalog := NewCatalog()
	catalog.CatalogComplexTypes(elementsSchema)
	catalog.CatalogSimpleTypes(typesSchema)

	ct, err := catalog.LookupComplexType("LoginInfo_T")
	assert.Nil(t, err)
	assert.Equal(t, "LoginInfo_T", ct.Name)
	assert.Equal(t, 7, len(ct.Sequence))

	for _, e := range ct.Sequence {
		fmt.Printf("ref: %s\n", e.Ref)
	}

	assert.Equal(t, false, catalog.IsSimpleType("foobar"))
	assert.Equal(t, true, catalog.IsSimpleType("AdminGroup_T"))
}
