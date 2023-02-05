/*
Package xml has some utilities for XML codec
*/
package xml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// StructToXML convert struct to XML
// The struct must have xml tags
//
//	:param interFace: The struct data with xml tags
//
// Example
//
//	var a NTP
//
//	thing, err := xml.StructToXML(a)
//
//	fmt.Println(thing)
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
//	:param interFace: A pointer to a variable of the struct data with xml tags
//	:param xmlData: The XML data
//
// Example
//
//	type NTP struct {
//		NTPServersData struct {
//			SourceInterface string `yaml:"source_interface" json:"source_interface" xml:"source_interface"`
//			NTPServerKey    []struct {
//				IPv4Host string `yaml:"ipv4_host" json:"ipv4_host" xml:"ipv4_host"`
//				Priority bool   `yaml:"priority" json:"priority" xml:"priority"`
//			} `yaml:"ntp_servers" json:"ntp_servers" xml:"ntp_servers"`
//		} `yaml:"ntp" json:"ntp" xml:"ntp"`
//	}
//
//	var a NTP
//
//	err := json.XMLToStruct(&a, ntpXml)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(a)
func XMLToStruct(interFace interface{}, xmlData string) error {
	if reflect.ValueOf(interFace).Kind() != reflect.Ptr {
		return errors.New("error XMLToStruct interFace must be a pointer")
	}

	stringReader := strings.NewReader(xmlData)
	decoder := xml.NewDecoder(stringReader)
	err := decoder.Decode(interFace)
	if err != nil {
		return errors.New(fmt.Sprintf("error XMLToStruct decoder.Decode %v", err))
	}

	return nil

}
