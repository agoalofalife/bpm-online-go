package bpm

import (
	"log"
	"github.com/andelf/go-curl"
	"os"
	"time"
	//"encoding/json"
	"encoding/json"
	//"fmt"
)
type auth struct{
	UserName string
	UserPassword string
}

func Bmp()  {
	url, _     :=  Config().String("auth.UrlLogin")
	login,_    :=  Config().String("auth.login")
	password,_ :=  Config().String("auth.Password")

	jsondi := &auth{
		UserName : login,
		UserPassword : password,
	}

	jsonString , _ := json.Marshal(jsondi)

	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		easy.Setopt(curl.OPT_URL, url)
		easy.Setopt(curl.OPT_VERBOSE, true)
		easy.Setopt(curl.OPT_COOKIEJAR, "./cookie.jar")
		//easy.Setopt(curl.OPT_POSTFIELDS, true)


		//form := curl.NewForm()
		//
		//form.Add("UserName", login) // your album id
		//form.AddFile("UserPassword", password)
		//form.Add("SolutionName", "TSBpm")
		//form.Add("TimeZoneOffset", "-120")
		//form.Add("Language", "Ru-ru")

		//easy.Setopt(curl.OPT_HTTPPOST, form)

		//postdata := "UserName=" + login + "&UserPassword=" + password
		easy.Setopt(curl.OPT_POSTFIELDS, jsonString)
	}
	if err := easy.Perform(); err != nil {
		println("ERROR: ", err.Error())
	} else{
		log.Println(easy.Perform())
	}


	time.Sleep(100000)
	os.Exit(2)
}
