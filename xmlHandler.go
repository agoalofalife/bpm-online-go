package bpm

import "errors"

type Xml struct {
	fields map[string]string
}


func XmlInit() *Xml {
	xml := Xml{}
	return &xml
}

func (xml Xml) Handler() (map[string]string, error) {
	maps := make(map[string]string)
	error := errors.New("")
	return maps, error
}
