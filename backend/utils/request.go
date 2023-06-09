package utils

import (
	"encoding/json"
	"io"
)

func GetRequestBody[T any](payload io.ReadCloser) (*T, error) {
	decodedPayload := new(T)
	decoder := json.NewDecoder(payload)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(decodedPayload)
	if err != nil {
		return nil, err
	}
	return decodedPayload, nil
}
