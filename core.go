package bpm

import (
	"log"
	//"os"
	"github.com/andelf/go-curl"

	"os"
)

func Bmp()  {
	url, _ := Config().String("auth.UrlLogin")

	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		easy.Setopt(curl.OPT_URL, url)
		easy.Setopt(curl.OPT_POSTFIELDS, true)
		easy.Perform()
	}

	log.Println(easy)
	os.Exit(2)
}
