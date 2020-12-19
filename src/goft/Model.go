package goft

import (
	"encoding/json"
	"log"
)

type Model interface {
	String() string
}

type Models string

func MakeModels(models interface{}) Models {
	b, err := json.Marshal(models)
	if err != nil {
		log.Println(err)
	}
	return Models(b)
}
