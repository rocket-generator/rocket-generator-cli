package data_mapper

import (
	"encoding/json"
	"io"
	"os"
)

type Mapper map[string]string

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

func MapString(mapper *Mapper, value string) string {
	if mapper == nil {
		return value
	}
	if mappedValue, ok := (*mapper)[value]; ok {
		return mappedValue
	}
	return value
}
