package bpm

import (
	"encoding/json"
	"github.com/andelf/go-curl"
	"io/ioutil"
	"os"
	"regexp"
)

type auth struct {
	UserName     string
	UserPassword string
}

type Cookie struct {
	fileCookie string
	prefixCSRF string
}

func AuthInit() Cookie {
	cookie := Cookie{}
	cookie.fileCookie = `./cookie.txt`
	cookie.prefixCSRF = `BPMCSRF`
	return cookie
}

func (c Cookie) GetCookie() (state bool, err string) {

	url, _ := Config().String("auth.UrlLogin")
	login, _ := Config().String("auth.login")
	password, _ := Config().String("auth.Password")

	jsonSchema := &auth{
		UserName:     login,
		UserPassword: password,
	}

	jsonString, _ := json.Marshal(jsonSchema)

	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		easy.Setopt(curl.OPT_URL, url)
		easy.Setopt(curl.OPT_VERBOSE, true)
		easy.Setopt(curl.OPT_COOKIEJAR, c.fileCookie)
		easy.Setopt(curl.OPT_HTTPHEADER, []string{"POST  HTTP/1.0", "Content-type: application/json"})
		easy.Setopt(curl.OPT_POSTFIELDS, string(jsonString))
	}
	if error := easy.Perform(); error != nil {
		err = error.Error()
	} else {
		state = true
	}

	return state, err
}

func (c Cookie) GetCsrf() string {
	if _, err := os.Stat(c.fileCookie); os.IsNotExist(err) {
		c.GetCookie()
	}
	fileContent, _ := ioutil.ReadFile(c.fileCookie)

	r, _ := regexp.Compile(`BPMCSRF\s(.+)`)
	matches := r.FindAllStringSubmatch(string(fileContent), 4)

	return matches[0][1]
}
