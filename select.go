package bpm

import (
	"github.com/andelf/go-curl"
	"log"
	"os"
	"strings"
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
	read.url = `?`
	read.core = core
	return &read
}

func (read *Select) Execute() interface{} {

	var page []byte
	escapeUrl := strings.Replace(read.url, " ", "%20", -1)
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
		easy.Setopt(curl.OPT_HTTPHEADER, []string{read.method + "  HTTP/1.0",
			"Content-type: " + read.core.handler.getContentType(),
			"Accept : " + read.core.handler.getAccept()})
	}

	if error := easy.Perform(); error != nil {
		log.Println(error)
		os.Exit(2)
	}

	code, _ := easy.Getinfo(curl.INFO_RESPONSE_CODE)
	isCode, refreshCookie := read.core.cookie.checkUnauthorized(code.(int))

	// if cookie obsolete example or they simply do not
	if isCode {
		refreshCookie()
		// repeat call
		return read.Execute()
	}

	resultHandler, _ := read.core.handler.Handler(page)
	return resultHandler
}

// filter constructor
func (read *Select) FilterConstructor(template string) (readyTemplate string) {
	readyTemplate = "$filter="
	readyTemplate += template
	return read.constructorUrl(readyTemplate)
}

/** Service resources can be obtained in the form of sort .
 * asc  ascending
 * desc descending
 * param howSort  asc | desc
 */
func (read Select) OrderBy(howSort ...string) string {

	var parameterSort string
	whoSort := howSort[0]

	if len(howSort) < 2 {
		parameterSort = `asc`
	} else {
		parameterSort = howSort[1]
	}

	param := `$orderby=`
	param += strings.ToUpper(whoSort[:1]) + whoSort[1:]
	param += " " + parameterSort

	return read.constructorUrl(param)
}

/** In bpm'online support the use of parameter $ the skip ,
 * which allows you to query the service resources ,
 * skipping the specified number of entries.
 */
func (read Select) Skip(number int) string {
	parameterQuery := "$skip=" + string(number)
	return read.constructorUrl(parameterQuery)
}

/** Restrictions in the sample query
 * If you want the request to return more than 40 records at a time, it can be implemented using the parameter $ top
 */
func (read Select) Amount(number int) string {
	parameterQuery := "$top=" + string(number)
	return read.constructorUrl(parameterQuery)
}

/** The number of records
 *  example SomeCollection/$count or SomeCollection/$count?$filter=...
 */
func (read Select) Count() string {
	read.url = "/"
	parameterQuery := "$count"
	return read.constructorUrl(parameterQuery)
}

/**
 * Contains guid
 * @param $guid
 * @return $this
 */
func (read Select) Guid(guid string) string {
	read.url = ""
	parameterQuery := "(guid"
	parameterQuery += "'"
	parameterQuery += guid
	parameterQuery += "'"
	parameterQuery += ")"
	return read.constructorUrl(parameterQuery)
}

// Get current string
func (read Select) GetUrl() string {
	return read.url
}

// function concat string URL
func (read *Select) constructorUrl(parameter string) string {

	if read.url == "?" {
		read.url += parameter
	} else if read.url == "" {
		read.url += parameter
	} else if read.url == "/" {
		read.url += parameter
	} else {
		read.url += "&" + parameter
	}
	return read.url
}
