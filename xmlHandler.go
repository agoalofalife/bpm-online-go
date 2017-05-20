package bpm

import "errors"

type Xml struct {
	fields map[string]string
	accept string
	contentType string
}


func XmlInit() DataType {
	xml := Xml{}
	xml.accept = "application/atom+xml;type=entry"
	xml.contentType = ""
	return &xml
}
// get content type in header request
func (xml Xml) getContentType() (contentType string)  {
	return xml.contentType
}

// get accept in header request
func (xml Xml) getAccept() (accept string) {
	return xml.accept
}

func (xml Xml) Handler() (map[string]string, error) {
	maps := make(map[string]string)
	error := errors.New("")
	return maps, error
}
