package bpm

import (
	"errors"
)
type Json struct {
	fields map[string]string
}


func JsonInit() DataType {
	json := Json{}
	return &json
}

func (json Json) Handler() (map[string]string, error) {
	maps := make(map[string]string)
	error := errors.New("")
	return maps, error
}