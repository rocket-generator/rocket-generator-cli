package b_parse_api_info_file

import (
	"encoding/json"
	"github.com/jinzhu/inflection"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
	"github.com/stoewer/go-strcase"
	"os"
	"strings"
)

type Process struct {
}

type API struct {
	Path          string   `json:"path"`
	Method        string   `json:"method"`
	Type          string   `json:"type"`
	SubType       string   `json:"subType"`
	Group         string   `json:"group"`
	RequireAuth   bool     `json:"requireAuth"`
	RequiredRoles []string `json:"requiredRoles"`
	TargetModel   string   `json:"targetModel"`
}

type APIs []API

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	if payload.ApiInfoFileName == "" {
		return payload, nil
	}
	apis, err := loadAPIInfoFile(payload.ApiInfoFileName)
	if err != nil {
		return nil, err
	}
	updateAPISpec(payload, apis)

	return payload, nil
}

func loadAPIInfoFile(filename string) (*APIs, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var apis APIs
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&apis); err != nil {
		return nil, err
	}

	for index, api := range apis {
		apis[index].Method = strcase.SnakeCase(api.Method)
		apis[index].Type = strcase.SnakeCase(api.Type)
	}
	return &apis, nil
}

func updateAPISpec(payload *newCommand.Payload, apis *APIs) {
	requestGroupMap := map[string]*objects.RequestGroup{}
	for _, api := range *apis {
		method := strcase.LowerCamelCase(api.Method)
		for index, request := range payload.OpenAPISpec.Requests {
			if request.Path == api.Path && request.Method.Snake == method {
				payload.OpenAPISpec.Requests[index].RequestType = strcase.LowerCamelCase(api.Type)
				payload.OpenAPISpec.Requests[index].RequestSubType = strcase.LowerCamelCase(api.SubType)
				if api.TargetModel != "" {
					apiTargetModelName := generateName(api.TargetModel)
					payload.OpenAPISpec.Requests[index].TargetModel = &apiTargetModelName
				}
				payload.OpenAPISpec.Requests[index].RequireAuth = api.RequireAuth
				payload.OpenAPISpec.Requests[index].RequiredRoles = api.RequiredRoles
				payload.OpenAPISpec.Requests[index].GroupRelativePath = generateGroupRelativePath(api.Group, request.Path)
				if requestGroupMap[api.Group] == nil {
					requestGroupMap[api.Group] = &objects.RequestGroup{
						PathPrefix: api.Group,
						Requests:   []*objects.Request{},
					}
				}
				requestGroupMap[api.Group].Requests = append(requestGroupMap[api.Group].Requests, request)
			}
		}
	}
	var requestGroups []*objects.RequestGroup
	for _, requestGroup := range requestGroupMap {
		requestGroups = append(requestGroups, requestGroup)
	}
	payload.OpenAPISpec.RequestGroups = &requestGroups
}

func generateGroupRelativePath(groupName string, path string) string {
	if path == "/" {
		return path
	}
	copiedPath := path

	if strings.HasPrefix(copiedPath, "/") {
		copiedPath = strings.TrimPrefix(copiedPath, "/")
	}

	if groupName == copiedPath {
		return "/"
	}

	if strings.HasPrefix(copiedPath, groupName+"/") {
		copiedPath = strings.TrimPrefix(copiedPath, groupName+"/")
	}

	return copiedPath
}

func generateName(name string) objects.Name {
	singular := inflection.Singular(name)
	plural := inflection.Plural(name)
	return objects.Name{
		Original: name,
		Default: objects.NameForm{
			Camel: strcase.LowerCamelCase(name),
			Title: strcase.UpperCamelCase(name),
			Snake: strcase.SnakeCase(name),
			Kebab: strcase.KebabCase(name),
		},
		Singular: objects.NameForm{
			Camel: strcase.LowerCamelCase(singular),
			Title: strcase.UpperCamelCase(singular),
			Snake: strcase.SnakeCase(singular),
			Kebab: strcase.KebabCase(singular),
		},
		Plural: objects.NameForm{
			Camel: strcase.LowerCamelCase(plural),
			Title: strcase.UpperCamelCase(plural),
			Snake: strcase.SnakeCase(plural),
			Kebab: strcase.KebabCase(plural),
		},
	}
}
