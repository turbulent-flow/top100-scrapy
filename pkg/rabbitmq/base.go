package rabbitmq

import (
	"encoding/json"
	"errors"
)

type arguments struct {
	Action string `json:"action,omitempty"`
	CategoryID int `json:"category_id,omitempty"`
	Page int `json:"page,omitempty"`
	Rank int `json:"rank,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}

func encode(obj interface{}) (encodedObj string, err error) {
	var encodedJson []byte
	encodedJson, err = json.Marshal(obj)
	if err != nil {
		content := "Failed to marshal an object"
		err = errors.New(content)
	}
	encodedObj = string(encodedJson)
	return encodedObj, err
}

func decode(data []byte, obj interface{}) error {
	return json.Unmarshal(data, obj)
}
