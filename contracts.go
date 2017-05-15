package bpm


type Action interface {
	Execute() bool
}

//type Handler interface {
//	Handler() (map[string]string , err error)
//}