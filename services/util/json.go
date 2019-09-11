package util

import (
	"encoding/json"
	"log"
)

// ToJSON marshal model to json string bytes
func ToJSON(model interface{}) []byte {
	str, err := json.Marshal(model)
	if err != nil {
		log.Printf("json marshal error %v", err)
	}
	return str
}

// ToModel unmarshal json string to model
func ToModel(jsonStr []byte, model interface{}) {
	err := json.Unmarshal(jsonStr, model)
	if err != nil {
		log.Printf("json unmarshal error %v", err)
	}
}