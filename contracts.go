package bpm

// Interface Action
// This is interface for actions type select, create, update and delete
type Action interface {
	Execute() bool
}

// Interface Handler
// Handler xml or json type in depends response bpm service
type DataType interface {
	Handler() (map[string]string, error)
}
