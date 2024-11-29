package data_mapper

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

type MapperElement map[string]string
type Mapper map[string]MapperElement

// Parse ...
func Parse(filePath string) (*Mapper, error) {
	jsonFile, err := os.Open(filePath)
	defer jsonFile.Close()

	if err != nil {
		return nil, err
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var mapper Mapper
	err = json.Unmarshal(byteValue, &mapper)
	if err != nil {
		return nil, err
	}

	return &mapper, nil
}

func MapString(mapper *Mapper, category string, value string) string {
	if mapper == nil {
		return value
	}
	lowerCaseCategory := strings.ToLower(category)
	if mapElement, ok := (*mapper)[lowerCaseCategory]; ok {
		lowerCaseValue := strings.ToLower(value)
		if mappedValue, ok := mapElement[lowerCaseValue]; ok {
			return mappedValue
		}
	}

	return value
}

func MapStringWithNil(mapper *Mapper, category string, value string) *string {
	if mapper == nil {
		return nil
	}
	lowerCaseCategory := strings.ToLower(category)
	if mapElement, ok := (*mapper)[lowerCaseCategory]; ok {
		lowerCaseValue := strings.ToLower(value)
		if mappedValue, ok := mapElement[lowerCaseValue]; ok {
			return &mappedValue
		}
	}

	return nil
}
