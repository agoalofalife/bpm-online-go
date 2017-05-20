package bpm

import (
	"errors"
	"github.com/clbanning/mxj"
	"fmt"
	"os"
	"log"
)

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

func (xml Xml) Handler(data []byte) (map[string]interface{}, error) {
	error := errors.New("")
	log.Println(data(string()), `test collection`)
	m, err := mxj.NewMapXml(data)
	v, _ := m.ValuesForKey("xml")
	//v, _ = m.ValuesForPath("feed.entry.content.properties")
	v, _ = m.ValuesForPath("service.*")
	//log.Println(v, `test collection`)
	os.Exit(2)
	if err != nil {
		error = errors.New("Error opening file")
	}

	maps := make(map[string]interface{})
	for _, vv := range v {
		for key, val := range vv.(map[string]interface{}) {
			maps[key] = val
			fmt.Println("\t\t", key, ":", val)
			os.Exit(2)
		}
		log.Println(xml)
		os.Exit(2)
	}
	log.Println(xml)
	os.Exit(2)

	return maps, error
}
