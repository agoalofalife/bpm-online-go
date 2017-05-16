package bpm

import (
	//"net/url"
	"github.com/andelf/go-curl"
	"log"
	"os"
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

	prepareUrl := read.core.collection + read.url
	//prepareUrl    := read.core.collection +  url.QueryEscape(read.url)
	urlHome, _ := Config().String("auth.UrlHome")
	urlHome += prepareUrl

	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		easy.Setopt(curl.OPT_URL, urlHome)
		easy.Setopt(curl.OPT_VERBOSE, true)
		easy.Setopt(curl.OPT_COOKIEJAR, read.core.cookie.fileCookie)
		easy.Setopt(curl.OPT_HTTPHEADER, []string{read.method + "  HTTP/1.0", "Content-type: application/json"})
	}
	if error := easy.Perform(); error != nil {
		log.Println(error)
		os.Exit(2)
	}
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
