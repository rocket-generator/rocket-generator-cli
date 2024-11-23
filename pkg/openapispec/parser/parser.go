package parser

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
	"github.com/stoewer/go-strcase"
	"net/url"
	"strings"
)

// Parse ...
func Parse(filePath string, namespace string, projectName string, organizationName string, typeMapper *data_mapper.Mapper) (*objects.API, error) {
	defaultRouteNamespace := namespace
	data := objects.API{
		FilePath:         filePath,
		BasePath:         "/",
		APINameSpace:     namespace,
		ProjectName:      projectName,
		OrganizationName: organizationName,
		Schemas:          map[string]*objects.Schema{},
		RouteNameSpace:   defaultRouteNamespace,
		RequestGroups:    nil,
	}
	openApiData, err := openapi3.NewLoader().LoadFromFile(filePath)
	if err != nil {
		return nil, err
	}
	if len(openApiData.Servers) > 0 {
		elements, err := url.Parse(openApiData.Servers[0].URL)
		if err != nil {
			return nil, err
		}
		data.BasePath = elements.Path
	}

	data.RouteNameSpace = buildRouteNameSpace(data.BasePath, defaultRouteNamespace)
	parseComponents(*openApiData.Components, &data, typeMapper)
	parsePaths(*openApiData.Paths, &data, typeMapper)

	return &data, nil
}

func buildRouteNameSpace(basePath string, defaultRouteNamespace string) string {
	if basePath == "/" || basePath == "" {
		return defaultRouteNamespace
	}

	elements := strings.Split(basePath, "/")
	name := ""
	for _, element := range elements {
		if element != "" {
			name = name + strcase.UpperCamelCase(element)
		}
	}
	return strcase.LowerCamelCase(name)
}
