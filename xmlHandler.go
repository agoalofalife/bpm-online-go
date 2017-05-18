package bpm

import "errors"

type Xml struct {

}
type XmlFile struct {
	Id        string `xml:"Id"`
	Number    string `xml:"Number"`
	S string `xml:"d>results>Id"`
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
