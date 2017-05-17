package bpm

type Core struct {
	collection string
	cookie     Cookie
	handler Handler
}

// Init application
func Start(collection string, handler string) *Core {
	core := Core{}
	core.collection = collection + "Collection"
	core.cookie	= AuthInit()
	if handler == "" {
		handler = XmlInit()
	}
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
