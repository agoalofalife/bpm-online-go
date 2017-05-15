package bpm

import (
	"log"
	"os"
)

type Core struct {
	collection string
}

// Init application
func Start(collection string) *Core {
	core := Core{}
	core.collection = collection
	return &core
}

func BmpTest() {
	temp := Read()
	temp.FilterConstructor("Id eq guid'68800b54-0e46-4388-b9f3-cbb45df42364'")
	log.Println(temp.Execute())
	os.Exit(2)
}
