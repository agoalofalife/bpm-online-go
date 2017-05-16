package bpm

import (
	"log"
	"os"
)

type Core struct {
	collection string
	Actions    map[string]Action
}

// Init application
func Start(collection string) *Core {
	mapsAction := make(map[string]Action)
	core := Core{}
	core.collection = collection + "Collection"
	mapsAction["read"] = Read()
	core.Actions = mapsAction
	return &core
}
// return link Read Action
func (core Core ) Read()  *Select {
	return Read()
}

func BmpTest() {
	temp := Read()
	temp.FilterConstructor("Id eq guid'68800b54-0e46-4388-b9f3-cbb45df42364'")
	log.Println(temp.Execute())
	os.Exit(2)
}
