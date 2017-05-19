package bpm

import "errors"

type Xml struct {

}
type XmlFile struct {
	S []string  `xml:",any"`
}
type XmlFeed struct {
	Ids []XmlContent `xml:"entry>content>properties"`
}
type XmlContent struct {
	Test string `xml:",any"`

	//Id string `xml:"Id"`/
	//Number string `xml:"Number"`
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
