package xsd_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/tmc/xsd"
)

var xsdPath = "./simple_schema.xsd"

func TestXSDParsing(t *testing.T) {
	content, err := ioutil.ReadFile(xsdPath)
	if err != nil {
		t.Fatal(err)
	}
	x, err := xsd.ParseXSD(bytes.NewReader(content))
	spew.Dump(x)
	return
	if err != nil {
		t.Fatal(err)
	}
	j, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
}
