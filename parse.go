package xsd

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

func ParseXSD(xsd io.Reader) (*Schema, error) {
	var value *Schema
	in, err := ioutil.ReadAll(xsd)
	if err != nil {
		return nil, err
	}
	if err := xml.Unmarshal(in, &value); err != nil {
		return nil, err
	}
	return value, nil
}
