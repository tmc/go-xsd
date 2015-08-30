package xsd

import (
	"encoding/xml"
	"io"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
)

func ParseXSD(xsd io.Reader) (*Schema, error) {
	var value *Schema
	in, err := ioutil.ReadAll(xsd)
	if err != nil {
		return nil, err
	}
	spew.Dump(in)
	if err := xml.Unmarshal(in, &value); err != nil {
		return nil, err
	}
	spew.Dump(value)
	return value, nil
}
