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
	catalog.CatalogSimpleTypes(elementsSchema)
	catalog.CatalogSimpleTypes(typesSchema)

	ct, err := catalog.LookupComplexType("LoginInfo_T")
	assert.Nil(t, err)
	assert.Equal(t, "LoginInfo_T", ct.Name)
	assert.Equal(t, 7, len(ct.Sequence))

	for _, e := range ct.Sequence {
		fmt.Printf("ref: %s\n", e.Ref)

		uq := UnqualifiedName(e.Ref)
		fmt.Printf("\tunqualfied name: %s", uq)

		elementDef, _ := catalog.LookupElementDef(uq)
		fmt.Printf("\ttype name: %s\n", elementDef.Type)

		et := catalog.ElementType(elementDef.Type)
		fmt.Printf("\ttype: %s\n", et)

		switch uq {
		case "securityFunctionBitmasks":
			assert.Equal(t, "Complex", catalog.ElementType(elementDef.Type).String())
		case "operatorProfileInfo":
			assert.Equal(t, "Complex", catalog.ElementType(elementDef.Type).String())
		case "securityFunctions":
			assert.Equal(t, "Complex", catalog.ElementType(elementDef.Type).String())
		case "securityToken":
			assert.Equal(t, "Simple", catalog.ElementType(elementDef.Type).String())
			assertGivenBaseType(t, catalog, elementDef.Type, "string")   
		case "sourceSystems":
			assert.Equal(t, "Complex", catalog.ElementType(elementDef.Type).String())
		case "daysToExpiration":
			assert.Equal(t, "Xsd", catalog.ElementType(elementDef.Type).String())
			assert.Equal(t, "int", catalog.XsdTypeToGolangType(elementDef.Type))
		case "timeOutMinutes":
			assert.Equal(t, "Simple", catalog.ElementType(elementDef.Type).String())
			assertGivenBaseType(t, catalog, elementDef.Type, "int")   
		}
	}

	assert.Equal(t, false, catalog.IsSimpleType("foobar"))
	assert.Equal(t, true, catalog.IsSimpleType("AdminGroup_T"))
}

func assertGivenBaseType(t *testing.T, c *Catalog, typeName string, assertedType string) {
	fmt.Println("checkin base type of ", typeName)
	base, err := c.XsdBaseType(typeName)
	assert.Nil(t,err)
	assert.Equal(t, base, assertedType)
}
