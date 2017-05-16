package bpm

import (
	"github.com/olebedev/config"
	"log"
)

func Config() *config.Config {

	cfg, err := config.ParseYamlFile("configuration.yml")

	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
