package main

import (
	"errors"
	"fmt"
)

type ElementType int

const (
	Xsd     ElementType = iota
	Simple  ElementType = iota
	Complex             = iota
)

func (et ElementType) String() string {
	switch et {
	case Xsd:
		return "Xsd"
	case Simple:
		return "Simple"
	case Complex:
		return "Complex"
	}
	return fmt.Sprintf("%d", et)
}

type Catalog struct {
	complexTypeMap map[string]ComplexType
	simpleTypeMap  map[string]SimpleType
	elementMap     map[string]Element
}

func NewCatalog() *Catalog {
	return &Catalog{
		make(map[string]ComplexType),
		make(map[string]SimpleType),
		make(map[string]Element),
	}
}

func (c *Catalog) catalogComplexType(ct ComplexType) {
	c.complexTypeMap[ct.Name] = ct
}

func (c *Catalog) catalogElement(e Element) {
	c.elementMap[e.Name] = e
}

func (c *Catalog) CatalogComplexTypes(s *Schema) {
	for _, ct := range s.ComplexTypes {
		fmt.Println("catalog ", ct.Name)
		c.catalogComplexType(ct)
	}

	for _, e := range s.Elements {
		fmt.Println("catalog element ", e.Name)
		c.catalogElement(e)
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

func (c *Catalog) catalogSimpleType(st SimpleType) {
	c.simpleTypeMap[st.Name] = st
}

func (c *Catalog) CatalogSimpleTypes(s *Schema) {
	for _, st := range s.SimpleTypes {
		fmt.Println("catalog ", st.Name)
		c.catalogSimpleType(st)
	}
}

func (c *Catalog) IsSimpleType(name string) bool {
	_, ok := c.simpleTypeMap[name]
	return ok
}

func (c *Catalog) LookupSimpleType(name string) (SimpleType, error) {
	st, ok := c.simpleTypeMap[name]
	if ok {
		return st, nil
	} else {
		return st, errors.New("Simple type not found")
	}
}

func (c *Catalog) LookupElementDef(name string) (Element, error) {
	e, ok := c.elementMap[name]
	if ok {
		return e, nil
	} else {
		return e, errors.New("Element def not found")
	}
}

func (c *Catalog) ElementType(typeName string) ElementType {
	if HasXsdPrefix(typeName) {
		return Xsd
	}

	uqTypeName := UnqualifiedName(typeName)
	if c.IsSimpleType(uqTypeName) {
		return Simple
	} else {
		return Complex
	}
}
