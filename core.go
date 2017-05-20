package bpm

import (
	"strings"
)

type Core struct {
	collection string
	cookie     Cookie
	handler    DataType
}

// Init application
func Start(param string) *Core {
	handlers := map[string]func() DataType{
		"xml":  XmlInit,
		"json": JsonInit,
	}
	core := Core{}
	split := strings.Split(param, ":")

	core.collection = split[0] + "Collection"
	core.cookie = AuthInit()
	core.handler = (handlers[split[1]])()
	return &core
}

// return link Read Action
func (core *Core) Read() *Select {
	return Read(core)
}

func BmpTest() {
	//temp.FilterConstructor("Id eq guid'68800b54-0e46-4388-b9f3-cbb45df42364'")
	//log.Println(temp.Execute())
	//os.Exit(2)
}
