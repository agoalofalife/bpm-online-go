package bpm

import (
	"strings"
	"github.com/andelf/go-curl"
	"log"
	"os"
	"fmt"
	"github.com/clbanning/mxj"
)

type Select struct {
	method string
	url    string
	core   *Core
}

// init action Read or Select
func Read(core *Core) *Select {
	read := Select{}
	read.method = `GET`
	read.core = core
	return &read
}

func (read *Select) Execute() bool {

	var page []byte
	escapeUrl  := strings.Replace(read.url, " ", "%20", -1)
	prepareUrl := read.core.collection + escapeUrl

	urlHome, _ := Config().String("auth.UrlHome")
	urlHome += prepareUrl

	easy := curl.EasyInit()
	defer easy.Cleanup()


	if easy != nil {
		easy.Setopt(curl.OPT_URL, urlHome)
		easy.Setopt(curl.OPT_VERBOSE, true)
		easy.Setopt(curl.OPT_COOKIEFILE, "./cookie.txt")
		easy.Setopt(curl.OPT_WRITEFUNCTION, func(ptr []byte, _ interface{}) bool {
			page = append(page, ptr...)
			return true
		})
		easy.Setopt(curl.OPT_NOPROGRESS, false)
		//easy.Setopt(curl.OPT_HTTPHEADER, []string{read.method + "  HTTP/1.0", "Content-type: application/json"})
	}
	if error := easy.Perform(); error != nil {
		log.Println(error)
		os.Exit(2)
	}
	m, err := mxj.NewMapXml(page)
	v, _ := m.ValuesForKey("xml")

	v, _ = m.ValuesForPath("feed.entry.content.properties")

	//var s XmlFeed
	//var s map[string]interface{}
	//err := xml.Unmarshal(page, &s)

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	 xml := make(map[string]interface{})
	for _, vv := range v {
		for key, val := range vv.(map[string]interface{}) {
			xml[key] = val
			fmt.Println("\t\t", key, ":", val)
			os.Exit(2)
		}
		log.Println(xml)
		os.Exit(2)
	}
	log.Println(xml)
	os.Exit(2)
	return true
}



// filter constructor
func (read *Select) FilterConstructor(template string) (readyTemplate string) {
	readyTemplate = "$filter="
	readyTemplate += template
	return read.constructorUrl(readyTemplate)
}

// Get current string
func (read Select) GetUrl() string {
	return read.url
}

// function concat string URL
func (read *Select) constructorUrl(parameter string) string {
	if read.url == "?" {
		read.url = parameter
	} else if read.url == "" {
		read.url += parameter
	} else if read.url == "/" {
		read.url += parameter
	} else {
		read.url += "&" + parameter
	}
	return read.url
}
