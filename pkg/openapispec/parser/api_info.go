package parser

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
	"github.com/stoewer/go-strcase"
)

func ParseAPIInfoFile(filename string, apiSpec *objects.API) error {
	apiInfo, err := loadAPIInfoFile(filename)
	if err != nil {
		return err
	}

	return updateAPISpec(apiSpec, apiInfo)
}

func loadAPIInfoFile(filename string) (*objects.APIInfos, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var apis objects.APIInfos
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

func updateAPISpec(apiSpec *objects.API, apis *objects.APIInfos) error {
	requestGroupMap := map[string]*objects.RequestGroup{}
	for _, api := range *apis {
		method := strcase.LowerCamelCase(api.Method)
		for index, request := range apiSpec.Requests {
			if request.Path == api.Path && request.Method.Snake == method {
				apiSpec.Requests[index].RequestType = strcase.LowerCamelCase(api.Type)
				apiSpec.Requests[index].RequestSubType = strcase.LowerCamelCase(api.SubType)
				if api.TargetModel != "" {
					apiTargetModelName := generateName(api.TargetModel)
					apiSpec.Requests[index].TargetModel = &apiTargetModelName
				}
				apiSpec.Requests[index].RequireAuth = api.RequireAuth
				apiSpec.Requests[index].RequiredRoles = api.RequiredRoles
				apiSpec.Requests[index].GroupRelativePath = generateGroupRelativePath(api.Group, request.Path)
				if len(api.AncestorModels) > 0 {
					for _, ancestorModel := range api.AncestorModels {
						ancestorModelName := generateName(ancestorModel.Name)
						parameterName := generateName(ancestorModel.Parameter)
						columnName := generateName(ancestorModel.Column)
						apiSpec.Requests[index].AncestorModels = append(apiSpec.Requests[index].AncestorModels, objects.AncestorModel{
							Name:      ancestorModelName,
							Parameter: parameterName,
							Column:    columnName,
						})
					}
				}
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
	apiSpec.RequestGroups = &requestGroups

	return nil
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
