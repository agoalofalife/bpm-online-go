package bpm

import (
	"errors"
)
type Json struct {
	fields map[string]string
	accept string
	contentType string
}


func JsonInit() DataType {
	json := Json{}
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
func (json Json) Handler() (map[string]string, error) {
	maps := make(map[string]string)
	error := errors.New("")
	return maps, error
}