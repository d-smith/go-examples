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
		fmt.Printf("ref: %s qualified: %v\n", e.Ref, IsQualifiedName(e.Ref))
	}

	assert.Equal(t, false, catalog.IsSimpleType("foobar"))
	assert.Equal(t, true, catalog.IsSimpleType("AdminGroup_T"))
	
	for _, e := range ct.Sequence {
		fmt.Printf("ref: %s\n", e.Ref)
		uq := UnqualifiedName(e.Ref)
		fmt.Printf("\tunqualfied name: %s", uq)
		elementDef, _ := catalog.LookupElementDef(uq)
		fmt.Printf("\ttype name: %s\n", elementDef.Type)
		fmt.Printf("\ttype: %s\n", func(typeName string, c *Catalog) string {
				if(HasXsdPrefix(typeName)) {
					return "xsd"
				} 
				
				uqTypeName := UnqualifiedName(typeName)
				if(c.IsSimpleType(uqTypeName)) {
					return "simple type"
				} else {
					return "complext type"
				}
			}(elementDef.Type, catalog))
	}
}
