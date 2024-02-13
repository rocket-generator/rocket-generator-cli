package b_parse_api_info_file

import (
	"encoding/json"
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
				payload.OpenAPISpec.Requests[index].RequestType = api.Type
				payload.OpenAPISpec.Requests[index].TargetModel = api.TargetModel
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
