package main

import (
	"errors"
	"fmt"
)

type Catalog struct {
	complexTypeMap map[string]ComplexType
	simpleTypeMap  map[string]SimpleType
}

func NewCatalog() *Catalog {
	return &Catalog{
		make(map[string]ComplexType),
		make(map[string]SimpleType),
	}
}

func (c *Catalog) catalogComplexType(ct ComplexType) {
	c.complexTypeMap[ct.Name] = ct
}

func (c *Catalog) CatalogComplexTypes(s *Schema) {
	for _, ct := range s.ComplexTypes {
		fmt.Println("catalog ", ct.Name)
		c.catalogComplexType(ct)
	}
}

func (c *Catalog) LookupComplexType(name string) (ComplexType, error) {
	ct, ok := c.complexTypeMap[name]
	if ok {
		return ct, nil
	} else {
		return ct, errors.New("Complex type not found")
	}
}
