/*
Package yaml has some utilities for YAML codec
*/
package yaml

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

// StructToYAML convert struct to YAML
// The struct must have yaml tags
//
//	:param interFace: The struct data with yaml tags
func StructToYAML(interFace interface{}) (yamlData string, err error) {
	var destinationString strings.Builder
	encoder := yaml.NewEncoder(&destinationString)
	err = encoder.Encode(&interFace)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error StructToYAML encoder.Encode %v", err))
	}

	return destinationString.String(), nil

}

// YAMLToStruct convert YAML to a Struct
// The struct must have yaml tags
//
//	:param interFace: The struct data with yaml tags
//	:param yamlData: The YAML data
func YAMLToStruct(interFace interface{}, yamlData string) (structData interface{}, err error) {
	stringReader := strings.NewReader(yamlData)
	decoder := yaml.NewDecoder(stringReader)
	err = decoder.Decode(&interFace)
	if err != nil {
		return interFace, errors.New(fmt.Sprintf("error YAMLToStruct decoder.Decode %v", err))
	}

	return interFace, nil

}
