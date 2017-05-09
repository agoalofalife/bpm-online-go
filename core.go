package bpm

import (
	"os"
	"log"
)


func Bmp()  {
	log.Println(	AuthInit().GetCsrf())
	os.Exit(2)
}


