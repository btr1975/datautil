/*
Package yaml has some utilities for YAML codec
*/
package yaml

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"reflect"
	"strings"
)

// StructToYAML convert struct to YAML
// The struct must have yaml tags
//
//	:param interFace: The struct data with yaml tags
//
// Example
//
//	var a NTP
//
//	thing, err := xml.StructToYAML(a)
//
//	fmt.Println(thing)
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
//	:param interFace: A pointer to a variable of the struct data with yaml tags
//	:param yamlData: The YAML data
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
//	err := json.YAMLToStruct(&a, ntpYaml)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(a)
func YAMLToStruct(interFace interface{}, yamlData string) error {
	if reflect.ValueOf(interFace).Kind() != reflect.Ptr {
		return errors.New("error YAMLToStruct interFace must be a pointer")
	}

	stringReader := strings.NewReader(yamlData)
	decoder := yaml.NewDecoder(stringReader)
	err := decoder.Decode(interFace)
	if err != nil {
		return errors.New(fmt.Sprintf("error YAMLToStruct decoder.Decode %v", err))
	}

	return nil

}
