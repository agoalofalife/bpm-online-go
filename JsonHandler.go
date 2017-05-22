package bpm

import (
	"errors"
	"github.com/clbanning/mxj"
	"log"
	"os"
)
type Json struct {
	fields map[string]string
	accept string
	contentType string
}


func JsonInit() DataType {
	json := Json{}
	json.accept      = "application/json;odata=verbose;"
	json.contentType = "application/json"
	return &json
}


// get content type for header request
func (json Json) getContentType() (contentType string)  {
	return  json.contentType
}

// get accept in header request
func (json Json) getAccept() (accept string) {
	return json.accept
}
// handler
func (json Json) Handler(data []byte) (interface{}, error) {
	error := errors.New("")

	m, err := mxj.NewMapJson(data)

	if err != nil {
		error = err
	}
	mxj.LeafUseDotNotation()
	l := m.LeafNodes()
	//leafvalues := m.LeafValues()
	log.Println(l)
	os.Exit(2)
	return  l, error
}