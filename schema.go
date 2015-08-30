package xsd

import (
	"encoding/xml"
	"fmt"

	"github.com/metaleap/go-xsd/types"
)

type Schema struct {
	ElemBase

	XMLName            xml.Name          `xml:"schema" json:"schema"`
	XMLNamespacePrefix string            `xml:"-" json:"-"`
	XMLNamespaces      map[string]string `xml:"-" json:"-"`
	XMLIncludedSchemas []*Schema         `xml:"-" json:"-"`
	XSDNamespacePrefix string            `xml:"-" json:"-"`
	XSDParentSchema    *Schema           `xml:"-" json:"-"`

	// attributes

	AttributeFormDefault string        `xml:"attributeFormDefault,attr" json:"attributeFormDefault,attr,omitempty"`
	BlockDefault         string        `xml:"blockDefault,attr" json:"blockDefault,attr,omitempty"`
	ElementFormDefault   string        `xml:"elementFormDefault,attr" json:"elementFormDefault,attr,omitempty"`
	FinalDefault         string        `xml:"finalDefault,attr" json:"finalDefault,attr,omitempty"`
	Lang                 xsdt.Language `xml:"lang,attr" json:"lang,attr,omitempty"`
	Id                   xsdt.Id       `xml:"id,attr" json:"id,attr,omitempty"`
	SchemaLocation       xsdt.AnyURI   `xml:"schemaLocation,attr" json:"schemaLocation,attr,omitempty"`
	TargetNamespace      xsdt.AnyURI   `xml:"targetNamespace,attr" json:"targetNamespace,attr,omitempty"`
	Version              xsdt.Token    `xml:"version,attr" json:"version,attr,omitempty"`

	// elements

	Annotation      *Annotation       `xml:"annotation" json:"annotation,omitempty"`
	Attributes      []*Attribute      `xml:"attribute" json:"attribute,omitempty"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup" json:"attributeGroup,omitempty"`
	ComplexTypes    []*ComplexType    `xml:"complexType" json:"complexType,omitempty"`
	Elements        []*Element        `xml:"element" json:"element,omitempty"`
	Groups          []*Group          `xml:"group" json:"group,omitempty"`
	Includes        []*Include        `xml:"include" json:"include,omitempty"`
	Imports         []*Import         `xml:"import" json:"import,omitempty"`
	Notations       []*Notation       `xml:"notation" json:"notation,omitempty"`
	Redefines       []*Redefine       `xml:"redefine" json:"redefine,omitempty"`
	SimpleTypes     []*SimpleType     `xml:"simpleType" json:"simpleType,omitempty"`

	loadUri string
}

func (s *Schema) allSchemas(loadedSchemas map[string]bool) (schemas []*Schema) {
	schemas = append(schemas, s)
	loadedSchemas[s.loadUri] = true
	for _, ss := range s.XMLIncludedSchemas {
		if v, ok := loadedSchemas[ss.loadUri]; ok && v {
			continue
		}
		schemas = append(schemas, ss.allSchemas(loadedSchemas)...)
	}
	return
}

func (s *Schema) RootSchema(pathSchemas []string) *Schema {
	if s.XSDParentSchema != nil {
		for _, sch := range pathSchemas {
			if s.XSDParentSchema.loadUri == sch {
				fmt.Printf("schema loop detected %+v - > %s!\n", pathSchemas, s.XSDParentSchema.loadUri)
				return s
			}
		}
		pathSchemas = append(pathSchemas, s.loadUri)
		return s.XSDParentSchema.RootSchema(pathSchemas)
	}
	return s
}
