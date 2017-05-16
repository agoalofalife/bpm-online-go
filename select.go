package bpm

import (
//"github.com/andelf/go-curl"
)

type Select struct {
	method string
	url    string
}

// init action Read or Select
func Read() *Select {
	read := Select{}
	read.method = `GET`
	return &read
}

func (read *Select) Execute() bool {
	//parameters := url.QueryEscape(read.url)
	//easy := curl.EasyInit()
	//defer easy.Cleanup()
	//if easy != nil {
	//	easy.Setopt(curl.OPT_URL, url)
	//	easy.Setopt(curl.OPT_VERBOSE, true)
	//	easy.Setopt(curl.OPT_COOKIEJAR, c.fileCookie)
	//	easy.Setopt(curl.OPT_HTTPHEADER, []string{"POST  HTTP/1.0", "Content-type: application/json"})
	//	easy.Setopt(curl.OPT_POSTFIELDS, string(jsonString))
	//}
	//if error := easy.Perform(); error != nil {
	//	err = error.Error()
	//} else {
	//	state = true
	//}
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
