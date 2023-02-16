package parser

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinzhu/inflection"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec"
	"github.com/stoewer/go-strcase"
	"strings"
)

func parsePaths(paths openapi3.Paths, data *openapispec.API) {

	for path, pathItem := range paths {
		for method, operation := range pathItem.Operations() {
			request := openapispec.Request{
				Path:             path,
				Method:           openapispec.NewNameForm(strings.ToUpper(method)),
				PathName:         generateName(getPathFormFromPath(path)),
				Description:      operation.Description,
				RouteNameSpace:   data.RouteNameSpace,
				OrganizationName: data.OrganizationName,
			}
			// Parameters
			for _, parameterReference := range operation.Parameters {
				parameter := parameterReference.Value
				request.Parameters = append(request.Parameters, &openapispec.Parameter{
					Name:     generateName(parameter.Name),
					In:       parameter.In,
					Required: parameter.Required,
				})
			}
			for _, parameterReference := range pathItem.Parameters {
				parameter := parameterReference.Value
				request.Parameters = append(request.Parameters, &openapispec.Parameter{
					Name:     generateName(parameter.Name),
					In:       parameter.In,
					Required: parameter.Required,
				})
			}
			if operation.RequestBody != nil {
				requestSchema := operation.RequestBody.Value.Content.Get("application/json")
				if requestSchema != nil {
					if requestSchema.Schema.Ref != "" {
						requestName := getSchemaNameFromSchema(requestSchema.Schema.Ref, requestSchema.Schema.Value)
						request.RequestSchemaName = generateName(requestName)
						request.RequestSchema = generateSchemaObject(requestSchema.Schema.Ref, requestSchema.Schema.Value)
					} else {
						requestSchemaName := strcase.SnakeCase(path) + "_" + strings.ToLower(method) + "_request"
						data.Schemas[requestSchemaName] = generateSchemaObject(requestSchemaName, requestSchema.Schema.Value)
						request.RequestSchemaName = generateName(requestSchemaName)
						request.RequestSchema = data.Schemas[requestSchemaName]
					}
				}
			}
			for statusCode, schemaObject := range operation.Responses {
				responseSchema := schemaObject.Value.Content.Get("application/json")

				if responseSchema != nil {
					responseName := getSchemaNameFromSchema(responseSchema.Schema.Ref, responseSchema.Schema.Value)
					schema, ok := data.Schemas[responseName]
					if ok {
						success := false
						if strings.HasPrefix(statusCode, "2") {
							success = true
						}
						response := &openapispec.Response{
							StatusCode: statusCode,
							Schema:     schema,
							Success:    success,
						}
						request.Responses = append(request.Responses, response)
						if success {
							request.SuccessResponse = response
						}
					}
				}
			}
			data.Requests = append(data.Requests, &request)
		}
	}
}

func getPathFormFromPath(path string) string {
	if path == "/" {
		return "index"
	}
	path = strings.TrimPrefix(path, "/")
	if strings.HasSuffix(path, "/") {
		path = path + "index"
	}

	elements := strings.Split(path, "/")
	resultElements := make([]string, 0)
	for _, element := range elements {
		if strings.HasPrefix(element, "{") && strings.HasSuffix(element, "}") {
			count := len(resultElements)
			if count > 0 {
				resultElements[count-1] = inflection.Singular(resultElements[count-1])
			}
		} else {
			resultElements = append(resultElements, strings.ToLower(element))
		}
	}
	pathName := strings.Join(resultElements, "_")
	return strings.ToLower(pathName)
}
