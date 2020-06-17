package rabbitmq

import (
	"encoding/json"
	"errors"
)

type content map[string]interface{}

func encode(obj *content) (encodedJson string, err error) {
	var jsonObj []byte
	jsonObj, err = json.Marshal(obj)
	if err != nil {
		c := "Failed to marshal an object"
		err = errors.New(c)
	}
	encodedJson = string(jsonObj)
	return encodedJson, err
}
