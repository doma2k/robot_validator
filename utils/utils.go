package utils

import (
	"encoding/json"
	"log"
)

func UnmarshalJSON(body []byte, data interface{}) {
	err := json.Unmarshal(body, data)
	if err != nil {
		log.Fatal(err)
	}
}

