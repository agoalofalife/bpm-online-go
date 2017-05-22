package bpm

import (
	"errors"
	"github.com/clbanning/mxj"
)


type Xml struct {
	fields      map[string]string
	accept      string
	contentType string
}

func XmlInit() DataType {
	xml := Xml{}
	xml.accept = "application/atom+xml;type=entry"
	xml.contentType = ""
	return &xml
}

// get content type in header request
func (xml Xml) getContentType() (contentType string) {
	return xml.contentType
}

// get accept in header request
func (xml Xml) getAccept() (accept string) {
	return xml.accept
}

func (xml Xml) Handler(data []byte) (interface{}, error) {
	error := errors.New("")

	m, err := mxj.NewMapXml(data)
	if err != nil {
		error = errors.New("Error data ")
	}
	v, _ := m.ValuesForKey("xml")
	v, _ = m.ValuesForPath("feed.entry.content.properties")
	if  len(v) == 0{
		v, _ = m.ValuesForPath("service.workspace.collection")
	}
	if len(v) == 0 {
		v, _ = m.ValuesForPath("error.*")
	}

	return v, error
}
