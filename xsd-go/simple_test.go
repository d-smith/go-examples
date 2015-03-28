package main

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleParse(t *testing.T) {
	simpleType := `
	<xsd:simpleType xmlns:xt="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/types" xmlns:xsd="http://www.w3.org/2001/XMLSchema" targetNamespace="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/types" elementFormDefault="qualified" attributeFormDefault="unqualified" name="AccessGroupValue_T">
		<xsd:annotation>
			<xsd:documentation>
				Type will be used to selectively retrive access group record  or super access group records or both
				</xsd:documentation>
		</xsd:annotation>
		<xsd:restriction base="xsd:string">
			<xsd:pattern value="(All|AccessGroup|SuperAccessGroup)"/>
		</xsd:restriction>
	</xsd:simpleType>
	`

	st := new(SimpleType)
	err := xml.Unmarshal([]byte(simpleType), &st)
	assert.Nil(t, err)
	assert.Equal(t, "AccessGroupValue_T", st.Name)
}

func TestSimpleInXSDParse(t *testing.T) {
	simpleType := `
	<xsd:schema xmlns:xt="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/types" xmlns:xsd="http://www.w3.org/2001/XMLSchema" targetNamespace="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/types" elementFormDefault="qualified" attributeFormDefault="unqualified">
	<xsd:simpleType name="AccessGroupValue_T">
		<xsd:annotation>
			<xsd:documentation>
				Type will be used to selectively retrive access group record  or super access group records or both
				</xsd:documentation>
		</xsd:annotation>
		<xsd:restriction base="xsd:string">
			<xsd:pattern value="(All|AccessGroup|SuperAccessGroup)"/>
		</xsd:restriction>
	</xsd:simpleType>
	</xsd:schema>
	`

	s := new(Schema)
	err := xml.Unmarshal([]byte(simpleType), &s)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(s.SimpleTypes))
	assert.Equal(t, "AccessGroupValue_T", s.SimpleTypes[0].Name)
}

func TestSimpleAndComplexInXSDParse(t *testing.T) {
	simpleType := `
	<xsd:schema xmlns:xt="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/types" xmlns:xsd="http://www.w3.org/2001/XMLSchema" targetNamespace="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/types" elementFormDefault="qualified" attributeFormDefault="unqualified">
	<xsd:simpleType name="AccessGroupValue_T">
		<xsd:annotation>
			<xsd:documentation>
				Type will be used to selectively retrive access group record  or super access group records or both
				</xsd:documentation>
		</xsd:annotation>
		<xsd:restriction base="xsd:string">
			<xsd:pattern value="(All|AccessGroup|SuperAccessGroup)"/>
		</xsd:restriction>
	</xsd:simpleType>
	<xsd:complexType name="MergeMapping_T">
		<xsd:annotation>
			<xsd:documentation>
        Contains summary of merge mapping and common admin details like name, description and history
        Usage:- Required for displaying already mapped local and remote fields in Create Merge Mapping screen.
      </xsd:documentation>
		</xsd:annotation>
		<xsd:sequence>
			<xsd:element ref="xt:mergeMapId" minOccurs="0"/>
			<xsd:element ref="xt:businessPartnerId"/>
			<xsd:element ref="xt:description" minOccurs="0"/>
			<xsd:element ref="xt:history" minOccurs="0"/>
			<xsd:element ref="xt:name"/>
			<xsd:element ref="xt:mergeMappingFieldSummaries"/>
		</xsd:sequence>
	</xsd:complexType>
	</xsd:schema>
	`

	s := new(Schema)
	err := xml.Unmarshal([]byte(simpleType), &s)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(s.SimpleTypes))
	assert.Equal(t, "AccessGroupValue_T", s.SimpleTypes[0].Name)
	assert.Equal(t, 1, len(s.ComplexTypes))
	assert.Equal(t, "MergeMapping_T", s.ComplexTypes[0].Name)
}
