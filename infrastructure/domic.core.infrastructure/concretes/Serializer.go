package concretes

import (
	"github.com/json-iterator/go"
)

type Serializer struct{}

func (serializer *Serializer) Serialize(object interface{}) (string, error) {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	content, err := json.MarshalToString(&object)

	if err != nil {
		return "", err
	}

	return content, nil

}

func (serializer *Serializer) Deserialize(stringifyPayload string, targetObject interface{}) error {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	err := json.UnmarshalFromString(stringifyPayload, &targetObject)

	return err

}

func NewSerializer() *Serializer {
	return &Serializer{}
}
