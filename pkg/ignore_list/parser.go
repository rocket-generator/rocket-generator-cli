package ignore_list

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Endpoint struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type rawIgnoreList struct {
	Endpoints []Endpoint `json:"endpoints"`
	Tables    []string   `json:"tables"`
	Responses []string   `json:"responses"`
}

type IgnoreList struct {
	Endpoints map[string]struct{} // key: "METHOD PATH"
	Tables    map[string]struct{} // key: table name
	Responses map[string]struct{} // key: response name
}

// Parse reads and parses the ignore list from a JSON file
func Parse(filePath string) (*IgnoreList, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var raw rawIgnoreList
	err = json.Unmarshal(byteValue, &raw)
	if err != nil {
		return nil, err
	}

	result := &IgnoreList{
		Endpoints: make(map[string]struct{}),
		Tables:    make(map[string]struct{}),
		Responses: make(map[string]struct{}),
	}

	for _, endpoint := range raw.Endpoints {
		// set lowercase
		key := strings.ToLower(endpoint.Method) + " " + strings.ToLower(endpoint.Path)
		result.Endpoints[key] = struct{}{}
		fmt.Println("Ignore Endpoint Entry: --> " + key)
	}

	for _, table := range raw.Tables {
		// set lowercase
		key := strings.ToLower(table)
		result.Tables[key] = struct{}{}
		fmt.Println("Ignore Table Entry: --> " + key)
	}

	for _, response := range raw.Responses {
		// set lowercase
		key := strings.ToLower(response)
		result.Responses[key] = struct{}{}
		fmt.Println("Ignore Response Entry: --> " + key)
	}

	return result, nil
}
