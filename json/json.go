/*
Package json has some utilities for JSON codec
*/
package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// StructToJSON convert struct to JSON
// The struct must have json tags
//
//	:param interFace: The struct data with json tags
func StructToJSON(interFace interface{}) (jsonData string, err error) {
	var destinationString strings.Builder
	encoder := json.NewEncoder(&destinationString)
	err = encoder.Encode(&interFace)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error StructToJSON encoder.Encode %v", err))
	}

	return destinationString.String(), nil

}

// JSONToStruct convert JSON to a Struct
// The struct must have json tags
//
//	:param interFace: The struct data with json tags
//	:param jsonData: The JSON data
func JSONToStruct(interFace interface{}, jsonData string) (structData interface{}, err error) {
	stringReader := strings.NewReader(jsonData)
	decoder := json.NewDecoder(stringReader)
	err = decoder.Decode(&interFace)
	if err != nil {
		return interFace, errors.New(fmt.Sprintf("error JSONToStruct decoder.Decode %v", err))
	}

	return interFace, nil

}
