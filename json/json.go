/*
Package json has some utilities for JSON codec
*/
package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// StructToJSON convert struct to JSON
// The struct must have json tags
//
//	:param interFace: The struct data with json tags
//
// Example
//
//	var a NTP
//
//	thing, err := xml.StructToJSON(a)
//
//	fmt.Println(thing)
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
//	:param interFace: A pointer to a variable of the struct data with json tags
//	:param jsonData: The JSON data
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
//	err := json.JSONToStruct(&a, ntpJson)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(a)
func JSONToStruct(interFace interface{}, jsonData string) error {
	if reflect.ValueOf(interFace).Kind() != reflect.Ptr {
		return errors.New("error JSONToStruct interFace must be a pointer")
	}

	stringReader := strings.NewReader(jsonData)
	decoder := json.NewDecoder(stringReader)
	err := decoder.Decode(interFace)
	if err != nil {
		return errors.New(fmt.Sprintf("error JSONToStruct decoder.Decode %v", err))
	}

	return nil

}
