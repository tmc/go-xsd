package xsd

import (
	xsdt "github.com/metaleap/go-xsd/types"
)

type XSDElement interface {
	ElementName() string
	Parent() XSDElement
}

type ElemBase struct {
	ElementParent  XSDElement  `json:"-"`
	XSDName        xsdt.NCName `json:"-"`
	XSDHasNameAttr bool        `json:"-"`
}

func (e *ElemBase) ElementName() string {
	if e.XSDName != "" {
		return string(e.XSDName)
	}
	return "UNKNOWN"
}

func (e *ElemBase) Parent() XSDElement { return e.ElementParent }

type All struct {
	ElemBase
	//	XMLName xml.Name `xml:"all" json:"all"`
	Id        xsdt.Id `xml:"id,attr" json:"id,attr,omitempty"`
	MaxOccurs string  `xml:"maxOccurs,attr" json:"maxOccurs,attr,omitempty"`
	MinOccurs string  `xml:"minOccurs,attr" json:"minOccurs,attr,omitempty"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
	Elements   []*Element  `xml:"element" json:"element,omitempty"`
}

type Annotation struct {
	ElemBase
	//	XMLName xml.Name `xml:"annotation" json:"annotation,omitempty"`
	AppInfos       []*AppInfo       `xml:"appinfo" json:"appinfo,omitempty"`
	Documentations []*Documentation `xml:"documentation" json:"documentation,omitempty"`
}

type Any struct {
	ElemBase
	//	XMLName xml.Name `xml:"any" json:"any,omitempty"`
	Id              xsdt.Id `xml:"id,attr" json:"id,attr,omitempty"`
	MaxOccurs       string  `xml:"maxOccurs,attr" json:"maxOccurs,attr,omitempty"`
	MinOccurs       string  `xml:"minOccurs,attr" json:"minOccurs,attr,omitempty"`
	Namespace       string  `xml:"namespace,attr" json:"namespace,attr,omitempty"`
	ProcessContents string  `xml:"processContents,attr" json:"processContents,attr,omitempty"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
}

type AnyAttribute struct {
	ElemBase
	//	XMLName xml.Name `xml:"anyAttribute" json:"anyAttribute,omitempty"`
	Id              xsdt.Id `xml:"id,attr" json:"id,attr,omitempty"`
	Namespace       string  `xml:"namespace,attr" json:"namespace,attr,omitempty"`
	ProcessContents string  `xml:"processContents,attr" json:"processContents,attr,omitempty"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
}

type AppInfo struct {
	ElemBase
	//	XMLName xml.Name `xml:"appinfo" json:"appinfo,omitempty"`
	Source xsdt.AnyURI `xml:"source,attr" json:"source,attr,omitempty"`
	CDATA  string      `xml:",chardata" json:",chardata,omitempty"`
}

type Attribute struct {
	ElemBase
	//	XMLName xml.Name `xml:"attribute" json:"attribute,omitempty"`
	Default string      `xml:"default,attr" json:"default,attr,omitempty"`
	Fixed   string      `xml:"fixed,attr" json:"fixed,attr,omitempty"`
	Form    string      `xml:"form,attr" json:"form,attr,omitempty"`
	Id      xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	Name    xsdt.NCName `xml:"name,attr" json:"name,attr,omitempty"`
	Ref     xsdt.Qname  `xml:"ref,attr" json:"ref,attr,omitempty"`
	Type    xsdt.Qname  `xml:"type,attr" json:"type,attr,omitempty"`
	Use     string      `xml:"use,attr" json:"use,attr,omitempty"`

	Annotation  *Annotation   `xml:"annotation" json:"annotation,omitempty"`
	SimpleTypes []*SimpleType `xml:"simpleType" json:"simpleType,omitempty"`
}

type AttributeGroup struct {
	ElemBase
	//	XMLName xml.Name `xml:"attributeGroup" json:"attributeGroup,omitempty"`
	Id   xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	Name xsdt.NCName `xml:"name,attr" json:"name,attr,omitempty"`
	Ref  xsdt.Qname  `xml:"ref,attr" json:"ref,attr,omitempty"`

	Annotation      *Annotation       `xml:"annotation" json:"annotation,omitempty"`
	AnyAttributes   []*AnyAttribute   `xml:"anyAttribute" json:"anyAttribute,omitempty"`
	Attributes      []*Attribute      `xml:"attribute" json:"attribute,omitempty"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup" json:"attributeGroup,omitempty"`
}

type Choice struct {
	ElemBase
	//	XMLName xml.Name `xml:"choice" json:"choice,omitempty"`
	Id        xsdt.Id `xml:"id,attr" json:"id,attr,omitempty"`
	MaxOccurs string  `xml:"maxOccurs,attr" json:"maxOccurs,attr,omitempty"`
	MinOccurs string  `xml:"minOccurs,attr" json:"minOccurs,attr,omitempty"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
	Anys       []*Any      `xml:"any" json:"any,omitempty"`
	Choices    []*Choice   `xml:"choice" json:"choice,omitempty"`
	Elements   []*Element  `xml:"element" json:"element,omitempty"`
	Groups     []*Group    `xml:"group" json:"group,omitempty"`
	Sequences  []*Sequence `xml:"sequence" json:"sequence,omitempty"`
}

type ComplexContent struct {
	ElemBase
	//	XMLName xml.Name `xml:"complexContent" json:"complexContent,omitempty"`
	Id    xsdt.Id `xml:"id,attr" json:"id,attr,omitempty"`
	Mixed bool    `xml:"mixed,attr" json:"mixed,attr,omitempty"`

	Annotation                *Annotation                `xml:"annotation" json:"annotation,omitempty"`
	ExtensionComplexContent   *ExtensionComplexContent   `xml:"extension" json:"extension,omitempty"`
	RestrictionComplexContent *RestrictionComplexContent `xml:"restriction" json:"restriction,omitempty"`
}

type ComplexType struct {
	ElemBase
	//	XMLName xml.Name `xml:"complexType" json:"complexType,omitempty"`
	Abstract bool        `xml:"abstract,attr" json:"abstract,attr,omitempty"`
	Block    string      `xml:"block,attr" json:"block,attr,omitempty"`
	Final    string      `xml:"final,attr" json:"final,attr,omitempty"`
	Id       xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	Mixed    bool        `xml:"mixed,attr" json:"mixed,attr,omitempty"`
	Name     xsdt.NCName `xml:"name,attr" json:"name,attr,omitempty"`

	All             *All              `xml:"all" json:"all,omitempty"`
	Annotation      *Annotation       `xml:"annotation" json:"annotation,omitempty"`
	AnyAttributes   []*AnyAttribute   `xml:"anyAttribute" json:"anyAttribute,omitempty"`
	Attributes      []*Attribute      `xml:"attribute" json:"attribute,omitempty"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup" json:"attributeGroup,omitempty"`
	Choice          *Choice           `xml:"choice" json:"choice,omitempty"`
	ComplexContent  *ComplexContent   `xml:"complexContent" json:"complexContent,omitempty"`
	Group           *Group            `xml:"group" json:"group,omitempty"`
	Sequence        *Sequence         `xml:"sequence" json:"sequence,omitempty"`
	SimpleContent   *SimpleContent    `xml:"simpleContent" json:"simpleContent,omitempty"`
}

type Documentation struct {
	ElemBase
	//	XMLName xml.Name `xml:"documentation" json:"documentation,omitempty"`
	Lang   xsdt.Language `xml:"lang,attr" json:"lang,attr,omitempty"`
	Source xsdt.AnyURI   `xml:"source,attr" json:"source,attr,omitempty"`
	CDATA  string        `xml:",chardata" json:",chardata,omitempty"`
}

type Element struct {
	ElemBase
	//	XMLName xml.Name `xml:"element" json:"element,omitempty"`
	Abstract          bool        `xml:"abstract,attr" json:"abstract,attr,omitempty"`
	Block             string      `xml:"block,attr" json:"block,attr,omitempty"`
	Default           string      `xml:"default,attr" json:"default,attr,omitempty"`
	Final             string      `xml:"final,attr" json:"final,attr,omitempty"`
	Fixed             string      `xml:"fixed,attr" json:"fixed,attr,omitempty"`
	Form              string      `xml:"form,attr" json:"form,attr,omitempty"`
	Id                xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	MaxOccurs         string      `xml:"maxOccurs,attr" json:"maxOccurs,attr,omitempty"`
	MinOccurs         string      `xml:"minOccurs,attr" json:"minOccurs,attr,omitempty"`
	Name              xsdt.NCName `xml:"name,attr" json:"name,attr,omitempty"`
	Nillable          bool        `xml:"nillable,attr" json:"nillable,attr,omitempty"`
	Ref               xsdt.Qname  `xml:"ref,attr" json:"ref,attr,omitempty"`
	SubstitutionGroup xsdt.Qname  `xml:"substitutionGroup,attr" json:"substitutionGroup,attr,omitempty"`
	Type              xsdt.Qname  `xml:"type,attr" json:"type,attr,omitempty"`

	Annotation  *Annotation   `xml:"annotation" json:"annotation,omitempty"`
	ComplexType *ComplexType  `xml:"complexType" json:"complexType,omitempty"`
	Keys        []*Key        `xml:"key" json:"key,omitempty"`
	KeyRef      *KeyRef       `xml:"keyref" json:"keyref,omitempty"`
	SimpleTypes []*SimpleType `xml:"simpleType" json:"simpleType,omitempty"`
	Unique      *Unique       `xml:"unique" json:"unique,omitempty"`
}

type ExtensionComplexContent struct {
	ElemBase
	//	XMLName xml.Name `xml:"extension" json:"extension,omitempty"`
	Base xsdt.Qname `xml:"base,attr" json:"base,attr,omitempty"`
	Id   xsdt.Id    `xml:"id,attr" json:"id,attr,omitempty"`

	All             *All              `xml:"all" json:"all,omitempty"`
	Annotation      *Annotation       `xml:"annotation" json:"annotation,omitempty"`
	AnyAttributes   []*AnyAttribute   `xml:"anyAttribute" json:"anyAttribute,omitempty"`
	Attributes      []*Attribute      `xml:"attribute" json:"attribute,omitempty"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup" json:"attributeGroup,omitempty"`
	Choices         []*Choice         `xml:"choice" json:"choice,omitempty"`
	Groups          []*Group          `xml:"group" json:"group,omitempty"`
	Sequences       []*Sequence       `xml:"sequence" json:"sequence,omitempty"`
}

type ExtensionSimpleContent struct {
	ElemBase
	//	XMLName xml.Name `xml:"extension" json:"extension,omitempty"`
	Base xsdt.Qname `xml:"base,attr" json:"base,attr,omitempty"`
	Id   xsdt.Id    `xml:"id,attr" json:"id,attr,omitempty"`

	Annotation      *Annotation       `xml:"annotation" json:"annotation,omitempty"`
	AnyAttributes   []*AnyAttribute   `xml:"anyAttribute" json:"anyAttribute,omitempty"`
	Attributes      []*Attribute      `xml:"attribute" json:"attribute,omitempty"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup" json:"attributeGroup,omitempty"`
}

type Field struct {
	ElemBase
	//	XMLName xml.Name `xml:"field" json:"field,omitempty"`
	Id    xsdt.Id `xml:"id,attr" json:"id,attr,omitempty"`
	Xpath string  `xml:"xpath,attr" json:"xpath,attr,omitempty"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
}

type Group struct {
	ElemBase
	//	XMLName xml.Name `xml:"group" json:"group,omitempty"`
	Id        xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	MaxOccurs string      `xml:"maxOccurs,attr" json:"maxOccurs,attr,omitempty"`
	MinOccurs string      `xml:"minOccurs,attr" json:"minOccurs,attr,omitempty"`
	Name      xsdt.NCName `xml:"name,attr" json:"name,attr,omitempty"`
	Ref       xsdt.Qname  `xml:"ref,attr" json:"ref,attr,omitempty"`

	All        *All        `xml:"all" json:"all,omitempty"`
	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
	Choice     *Choice     `xml:"choice" json:"choice,omitempty"`
	Sequence   *Sequence   `xml:"sequence" json:"sequence,omitempty"`
}

type Include struct {
	ElemBase
	//	XMLName xml.Name `xml:"include" json:"include,omitempty"`
	Id             xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	SchemaLocation xsdt.AnyURI `xml:"schemaLocation,attr" json:"schemaLocation,attr,omitempty"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
}

type Import struct {
	ElemBase
	//	XMLName xml.Name `xml:"import" json:"import,omitempty"`
	Id             xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	Namespace      string      `xml:"namespace,attr" json:"namespace,attr,omitempty"`
	SchemaLocation xsdt.AnyURI `xml:"schemaLocation,attr" json:"schemaLocation,attr,omitempty"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
}

type Key struct {
	ElemBase
	//	XMLName xml.Name `xml:"key" json:"key,omitempty"`
	Id   xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	Name xsdt.NCName `xml:"name,attr" json:"name,attr,omitempty"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
	Field      *Field      `xml:"field" json:"field,omitempty"`
	Selector   *Selector   `xml:"selector" json:"selector,omitempty"`
}

type KeyRef struct {
	ElemBase
	//	XMLName xml.Name `xml:"keyref" json:"keyref,omitempty"`
	Id    xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	Name  xsdt.NCName `xml:"name,attr" json:"name,attr,omitempty"`
	Refer xsdt.Qname  `xml:"refer,attr" json:"refer,attr,omitempty"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
	Field      *Field      `xml:"field" json:"field,omitempty"`
	Selector   *Selector   `xml:"selector" json:"selector,omitempty"`
}

type List struct {
	ElemBase
	//	XMLName xml.Name `xml:"list" json:"list,omitempty"`
	Id       xsdt.Id    `xml:"id,attr" json:"id,attr,omitempty"`
	ItemType xsdt.Qname `xml:"itemType,attr" json:"itemType,attr,omitempty"`

	Annotation  *Annotation   `xml:"annotation" json:"annotation,omitempty"`
	SimpleTypes []*SimpleType `xml:"simpleType" json:"simpleType,omitempty"`
}

type Notation struct {
	ElemBase
	//	XMLName xml.Name `xml:"notation" json:"notation,omitempty"`
	Id     xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	Name   xsdt.NCName `xml:"name,attr" json:"name,attr,omitempty"`
	Public string      `xml:"public,attr" json:"public,attr,omitempty"`
	System xsdt.AnyURI `xml:"system,attr" json:"system,attr,omitempty"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
}

type Redefine struct {
	ElemBase
	//	XMLName xml.Name `xml:"redefine" json:"redefine,omitempty"`
	Id             xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	SchemaLocation xsdt.AnyURI `xml:"schemaLocation,attr" json:"schemaLocation,attr,omitempty"`

	Annotation      *Annotation       `xml:"annotation" json:"annotation,omitempty"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup" json:"attributeGroup,omitempty"`
	ComplexTypes    []*ComplexType    `xml:"complexType" json:"complexType,omitempty"`
	Groups          []*Group          `xml:"group" json:"group,omitempty"`
	SimpleTypes     []*SimpleType     `xml:"simpleType" json:"simpleType,omitempty"`
}

type RestrictionComplexContent struct {
	ElemBase
	//	XMLName xml.Name `xml:"restriction" json:"restriction,omitempty"`
	Base xsdt.Qname `xml:"base,attr" json:"base,attr,omitempty"`
	Id   xsdt.Id    `xml:"id,attr" json:"id,attr,omitempty"`

	All             *All              `xml:"all" json:"all,omitempty"`
	Annotation      *Annotation       `xml:"annotation" json:"annotation,omitempty"`
	AnyAttributes   []*AnyAttribute   `xml:"anyAttribute" json:"anyAttribute,omitempty"`
	Attributes      []*Attribute      `xml:"attribute" json:"attribute,omitempty"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup" json:"attributeGroup,omitempty"`
	Choices         []*Choice         `xml:"choice" json:"choice,omitempty"`
	Sequences       []*Sequence       `xml:"sequence" json:"sequence,omitempty"`
}

type RestrictionSimpleContent struct {
	ElemBase
	//	XMLName xml.Name `xml:"restriction" json:"restriction,omitempty"`
	Base xsdt.Qname `xml:"base,attr" json:"base,attr,omitempty"`
	Id   xsdt.Id    `xml:"id,attr" json:"id,attr,omitempty"`

	Annotation      *Annotation                      `xml:"annotation" json:"annotation,omitempty"`
	AnyAttributes   []*AnyAttribute                  `xml:"anyAttribute" json:"anyAttribute,omitempty"`
	Attributes      []*Attribute                     `xml:"attribute" json:"attribute,omitempty"`
	AttributeGroups []*AttributeGroup                `xml:"attributeGroup" json:"attributeGroup,omitempty"`
	Enumerations    []*RestrictionSimpleEnumeration  `xml:"enumeration" json:"enumeration,omitempty"`
	FractionDigits  *RestrictionSimpleFractionDigits `xml:"fractionDigits" json:"fractionDigits,omitempty"`
	Length          *RestrictionSimpleLength         `xml:"length" json:"length,omitempty"`
	MaxExclusive    *RestrictionSimpleMaxExclusive   `xml:"maxExclusive" json:"maxExclusive,omitempty"`
	MaxInclusive    *RestrictionSimpleMaxInclusive   `xml:"maxInclusive" json:"maxInclusive,omitempty"`
	MaxLength       *RestrictionSimpleMaxLength      `xml:"maxLength" json:"maxLength,omitempty"`
	MinExclusive    *RestrictionSimpleMinExclusive   `xml:"minExclusive" json:"minExclusive,omitempty"`
	MinInclusive    *RestrictionSimpleMinInclusive   `xml:"minInclusive" json:"minInclusive,omitempty"`
	MinLength       *RestrictionSimpleMinLength      `xml:"minLength" json:"minLength,omitempty"`
	Pattern         *RestrictionSimplePattern        `xml:"pattern" json:"pattern,omitempty"`
	SimpleTypes     []*SimpleType                    `xml:"simpleType" json:"simpleType,omitempty"`
	TotalDigits     *RestrictionSimpleTotalDigits    `xml:"totalDigits" json:"totalDigits,omitempty"`
	WhiteSpace      *RestrictionSimpleWhiteSpace     `xml:"whiteSpace" json:"whiteSpace,omitempty"`
}

type RestrictionSimpleEnumeration struct {
	ElemBase
	//	XMLName xml.Name `xml:"enumeration" json:"enumeration,omitempty"`
	Value string `xml:"value,attr" json:"value,attr,omitempty"`
}

type RestrictionSimpleFractionDigits struct {
	ElemBase
	//	XMLName xml.Name `xml:"fractionDigits" json:"fractionDigits,omitempty"`
	Value string `xml:"value,attr" json:"value,attr,omitempty"`
}

type RestrictionSimpleLength struct {
	ElemBase
	//	XMLName xml.Name `xml:"length" json:"length,omitempty"`
	Value string `xml:"value,attr" json:"value,attr,omitempty"`
}

type RestrictionSimpleMaxExclusive struct {
	ElemBase
	//	XMLName xml.Name `xml:"maxExclusive" json:"maxExclusive,omitempty"`
	Value string `xml:"value,attr" json:"value,attr,omitempty"`
}

type RestrictionSimpleMaxInclusive struct {
	ElemBase
	//	XMLName xml.Name `xml:"maxInclusive" json:"maxInclusive,omitempty"`
	Value string `xml:"value,attr" json:"value,attr,omitempty"`
}

type RestrictionSimpleMaxLength struct {
	ElemBase
	//	XMLName xml.Name `xml:"maxLength" json:"maxLength,omitempty"`
	Value string `xml:"value,attr" json:"value,attr,omitempty"`
}

type RestrictionSimpleMinExclusive struct {
	ElemBase
	//	XMLName xml.Name `xml:"minExclusive" json:"minExclusive,omitempty"`
	Value string `xml:"value,attr" json:"value,attr,omitempty"`
}

type RestrictionSimpleMinInclusive struct {
	ElemBase
	//	XMLName xml.Name `xml:"minInclusive" json:"minInclusive,omitempty"`
	Value string `xml:"value,attr" json:"value,attr,omitempty"`
}

type RestrictionSimpleMinLength struct {
	ElemBase
	//	XMLName xml.Name `xml:"minLength" json:"minLength,omitempty"`
	Value string `xml:"value,attr" json:"value,attr,omitempty"`
}

type RestrictionSimplePattern struct {
	ElemBase
	//	XMLName xml.Name `xml:"pattern" json:"pattern,omitempty"`
	Value string `xml:"value,attr" json:"value,attr,omitempty"`
}

type RestrictionSimpleTotalDigits struct {
	ElemBase
	//	XMLName xml.Name `xml:"totalDigits" json:"totalDigits,omitempty"`
	Value string `xml:"value,attr" json:"value,attr,omitempty"`
}

type RestrictionSimpleType struct {
	ElemBase
	//	XMLName xml.Name `xml:"restriction" json:"restriction,omitempty"`
	Base xsdt.Qname `xml:"base,attr" json:"base,attr,omitempty"`
	Id   xsdt.Id    `xml:"id,attr" json:"id,attr,omitempty"`

	Annotation     *Annotation                      `xml:"annotation" json:"annotation,omitempty"`
	Enumerations   []*RestrictionSimpleEnumeration  `xml:"enumeration" json:"enumeration,omitempty"`
	FractionDigits *RestrictionSimpleFractionDigits `xml:"fractionDigits" json:"fractionDigits,omitempty"`
	Length         *RestrictionSimpleLength         `xml:"length" json:"length,omitempty"`
	MaxExclusive   *RestrictionSimpleMaxExclusive   `xml:"maxExclusive" json:"maxExclusive,omitempty"`
	MaxInclusive   *RestrictionSimpleMaxInclusive   `xml:"maxInclusive" json:"maxInclusive,omitempty"`
	MaxLength      *RestrictionSimpleMaxLength      `xml:"maxLength" json:"maxLength,omitempty"`
	MinExclusive   *RestrictionSimpleMinExclusive   `xml:"minExclusive" json:"minExclusive,omitempty"`
	MinInclusive   *RestrictionSimpleMinInclusive   `xml:"minInclusive" json:"minInclusive,omitempty"`
	MinLength      *RestrictionSimpleMinLength      `xml:"minLength" json:"minLength,omitempty"`
	Pattern        *RestrictionSimplePattern        `xml:"pattern" json:"pattern,omitempty"`
	SimpleTypes    []*SimpleType                    `xml:"simpleType" json:"simpleType,omitempty"`
	TotalDigits    *RestrictionSimpleTotalDigits    `xml:"totalDigits" json:"totalDigits,omitempty"`
	WhiteSpace     *RestrictionSimpleWhiteSpace     `xml:"whiteSpace" json:"whiteSpace,omitempty"`
}

type RestrictionSimpleWhiteSpace struct {
	ElemBase
	//	XMLName xml.Name `xml:"whiteSpace" json:"whiteSpace,omitempty"`
	Value string `xml:"value,attr" json:"value,attr,omitempty"`
}

type Selector struct {
	ElemBase
	//	XMLName xml.Name `xml:"selector" json:"selector,omitempty"`
	Id    xsdt.Id `xml:"id,attr" json:"id,attr,omitempty"`
	Xpath string  `xml:"xpath,attr" json:"xpath,attr,omitempty"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
}

type Sequence struct {
	ElemBase
	//	XMLName xml.Name `xml:"sequence" json:"sequence,omitempty"`
	Id        xsdt.Id `xml:"id,attr" json:"id,attr,omitempty"`
	MaxOccurs string  `xml:"maxOccurs,attr" json:"maxOccurs,attr,omitempty"`
	MinOccurs string  `xml:"minOccurs,attr" json:"minOccurs,attr,omitempty"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
	Anys       []*Any      `xml:"any" json:"any,omitempty"`
	Choices    []*Choice   `xml:"choice" json:"choice,omitempty"`
	Elements   []*Element  `xml:"element" json:"element,omitempty"`
	Groups     []*Group    `xml:"group" json:"group,omitempty"`
	Sequences  []*Sequence `xml:"sequence" json:"sequence,omitempty"`
}

type SimpleContent struct {
	ElemBase
	//	XMLName xml.Name `xml:"simpleContent" json:"simpleContent,omitempty"`
	Id xsdt.Id `xml:"id,attr" json:"id,attr,omitempty"`

	Annotation               *Annotation               `xml:"annotation" json:"annotation,omitempty"`
	ExtensionSimpleContent   *ExtensionSimpleContent   `xml:"extension" json:"extension,omitempty"`
	RestrictionSimpleContent *RestrictionSimpleContent `xml:"restriction" json:"restriction,omitempty"`
}

type SimpleType struct {
	ElemBase
	//	XMLName xml.Name `xml:"simpleType" json:"simpleType,omitempty"`
	Final string      `xml:"final,attr" json:"final,attr,omitempty"`
	Id    xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	Name  xsdt.NCName `xml:"name,attr" json:"name,attr,omitempty"`

	Annotation            *Annotation            `xml:"annotation" json:"annotation,omitempty"`
	List                  *List                  `xml:"list" json:"list,omitempty"`
	RestrictionSimpleType *RestrictionSimpleType `xml:"restriction" json:"restriction,omitempty"`
	Union                 *Union                 `xml:"union" json:"union,omitempty"`
}

type Union struct {
	ElemBase
	//	XMLName xml.Name `xml:"union" json:"union,omitempty"`
	Id          xsdt.Id `xml:"id,attr" json:"id,attr,omitempty"`
	MemberTypes string  `xml:"memberTypes,attr" json:"memberTypes,attr,omitempty"`

	Annotation  *Annotation   `xml:"annotation" json:"annotation,omitempty"`
	SimpleTypes []*SimpleType `xml:"simpleType" json:"simpleType,omitempty"`
}

type Unique struct {
	ElemBase
	//	XMLName xml.Name `xml:"unique" json:"unique,omitempty"`
	Id   xsdt.Id     `xml:"id,attr" json:"id,attr,omitempty"`
	Name xsdt.NCName `xml:"name,attr"`

	Annotation *Annotation `xml:"annotation" json:"annotation,omitempty"`
	Field      *Field      `xml:"field" json:"field,omitempty"`
	Selector   *Selector   `xml:"selector" json:"selector,omitempty"`
}

func Flattened(choices []*Choice, seqs []*Sequence) (allChoices []*Choice, allSeqs []*Sequence) {
	var tmpChoices []*Choice
	var tmpSeqs []*Sequence
	for _, ch := range choices {
		if ch != nil {
			allChoices = append(allChoices, ch)
			tmpChoices, tmpSeqs = Flattened(ch.Choices, ch.Sequences)
			allChoices = append(allChoices, tmpChoices...)
			allSeqs = append(allSeqs, tmpSeqs...)
		}
	}
	for _, seq := range seqs {
		if seq != nil {
			allSeqs = append(allSeqs, seq)
			tmpChoices, tmpSeqs = Flattened(seq.Choices, seq.Sequences)
			allChoices = append(allChoices, tmpChoices...)
			allSeqs = append(allSeqs, tmpSeqs...)
		}
	}
	return
}
