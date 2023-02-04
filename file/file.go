/*
Package file is a group of utilities for reading and writing files
*/
package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// yamlExtensionsAllowed holds allowed YAML file extensions
var yamlExtensionsAllowed = []string{
	"yml",
	"yaml",
}

// jsonExtensionsAllowed holds allowed JSON file extensions
var jsonExtensionsAllowed = []string{
	"json",
}

// xmlExtensionsAllowed holds allowed XML file extensions
var xmlExtensionsAllowed = []string{
	"xml",
}

// FileExtensionChecker Checks file name extensions
//
//	:param fileName: The file name including the extension
//	:param allowedExtensions: The allowed extensions
func FileExtensionChecker(fileName string, allowedExtensions []string) (err error) {
	fileNameSplit := strings.Split(fileName, ".")

	extensionFound := false
	for _, extension := range allowedExtensions {
		if fileNameSplit[1] == extension {
			extensionFound = true
			break
		}
	}

	if !extensionFound {
		return errors.New(fmt.Sprintf("error FileExtensionChecker %v bad extenstion, "+
			"allowed is %v", fileName, allowedExtensions))
	}

	return nil
}

// ReadYAMLFile read a YAML file, allowed extensions are yml, and yaml
//
//	:param path: The full path to the file
func ReadYAMLFile(path string) (yamlData string, err error) {
	pathSplit := strings.Split(path, "/")

	fileName := pathSplit[len(pathSplit)-1]

	err = FileExtensionChecker(fileName, yamlExtensionsAllowed)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error ReadYAMLFile %v", err))
	}

	data, err := ReadFile(path)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error ReadYAMLFile %v", err))
	}
	return data, nil
}

// ReadJSONFile read a JSON file, allowed extension is json
//
//	:param path: The full path to the file
func ReadJSONFile(path string) (jsonData string, err error) {
	pathSplit := strings.Split(path, "/")

	fileName := pathSplit[len(pathSplit)-1]

	err = FileExtensionChecker(fileName, jsonExtensionsAllowed)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error ReadJSONFile %v", err))
	}

	data, err := ReadFile(path)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error ReadJSONFile %v", err))
	}
	return data, nil
}

// ReadXMLFile read an XML file, allowed extension is xml
//
//	:param path: The full path to the file
func ReadXMLFile(path string) (jsonData string, err error) {
	pathSplit := strings.Split(path, "/")

	fileName := pathSplit[len(pathSplit)-1]

	err = FileExtensionChecker(fileName, xmlExtensionsAllowed)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error ReadXMLFile %v", err))
	}

	data, err := ReadFile(path)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error ReadXMLFile %v", err))
	}
	return data, nil
}

// ReadFile read data from a file
//
//	:param path: The full path to the file
func ReadFile(path string) (data string, err error) {
	src, err := os.ReadFile(path)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error ReadFile os.ReadFile %v", err))
	}
	return string(src), nil
}

// WriteYAMLFile write a YAML file, allowed extensions are yml, and yaml
//
//	:param path: The full path to the file
//	:param dat: The YAML data to write
func WriteYAMLFile(path string, data string) (err error) {
	pathSplit := strings.Split(path, "/")

	fileName := pathSplit[len(pathSplit)-1]

	err = FileExtensionChecker(fileName, yamlExtensionsAllowed)

	if err != nil {
		return errors.New(fmt.Sprintf("error WriteYAMLFile %v", err))
	}

	err = WriteFile(path, data)

	if err != nil {
		return errors.New(fmt.Sprintf("error WriteYAMLFile %v", err))
	}
	return nil
}

// WriteJSONFile write a JSON file, allowed extensions are json
//
//	:param path: The full path to the file
//	:param dat: The JSON data to write
func WriteJSONFile(path string, data string) (err error) {
	pathSplit := strings.Split(path, "/")

	fileName := pathSplit[len(pathSplit)-1]

	err = FileExtensionChecker(fileName, jsonExtensionsAllowed)

	if err != nil {
		return errors.New(fmt.Sprintf("error WriteJSONFile %v", err))
	}

	err = WriteFile(path, data)

	if err != nil {
		return errors.New(fmt.Sprintf("error WriteJSONFile %v", err))
	}
	return nil
}

// WriteXMLFile write an XML file, allowed extensions are xml
//
//	:param path: The full path to the file
//	:param dat: The XML data to write
func WriteXMLFile(path string, data string) (err error) {
	pathSplit := strings.Split(path, "/")

	fileName := pathSplit[len(pathSplit)-1]

	err = FileExtensionChecker(fileName, xmlExtensionsAllowed)

	if err != nil {
		return errors.New(fmt.Sprintf("error WriteXMLFile %v", err))
	}

	err = WriteFile(path, data)

	if err != nil {
		return errors.New(fmt.Sprintf("error WriteXMLFile %v", err))
	}
	return nil
}

// WriteFile write data to a file
//
//	:param path: The full path to the file
//	:param data: The string data you want to write
func WriteFile(path string, data string) (err error) {
	err = os.WriteFile(path, []byte(data), 0644)

	if err != nil {
		return errors.New(fmt.Sprintf("error WriteFile %v", err))
	}

	return err
}
