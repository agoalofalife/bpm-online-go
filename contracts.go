package bpm

// Interface Action
// This is interface for actions type select, create, update and delete
type Action interface {
	Execute() bool
}

// Interface Handler
// Handler xml or json type in depends response bpm service
type DataType interface {
	Handler(data []byte) (map[string]interface{}, error)
	getAccept() string
	getContentType() string
}

// turn  this is implements turn
type turn interface {
	// add new process
	add(process turn) (bool)
	// remove process
	remove(process turn)(bool)
}

