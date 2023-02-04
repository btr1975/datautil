/*
Package xml has some utilities for XML codec
*/
package xml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

// StructToXML convert struct to XML
// The struct must have xml tags
//
//	:param interFace: The struct data with xml tags
func StructToXML(interFace interface{}) (xmlData string, err error) {
	var destinationString strings.Builder
	encoder := xml.NewEncoder(&destinationString)
	err = encoder.Encode(&interFace)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error StructToXML encoder.Encode %v", err))
	}

	return destinationString.String(), nil

}

// XMLToStruct convert XML to a Struct
// The struct must have xml tags
//
//	:param interFace: The struct data with xml tags
//	:param xmlData: The XML data
func XMLToStruct(interFace interface{}, xmlData string) (structData interface{}, err error) {
	stringReader := strings.NewReader(xmlData)
	decoder := xml.NewDecoder(stringReader)
	err = decoder.Decode(&interFace)
	if err != nil {
		return interFace, errors.New(fmt.Sprintf("error XMLToStruct decoder.Decode %v", err))
	}

	return interFace, nil

}
