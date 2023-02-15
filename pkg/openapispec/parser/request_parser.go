package parser

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec"
	"github.com/stoewer/go-strcase"
	"strings"
)

func parsePaths(paths openapi3.Paths, data *openapispec.API) {

	for path, pathItem := range paths {
		for method, operation := range pathItem.Operations() {
			request := openapispec.Request{
				Path:           path,
				Method:         openapispec.NewNameForm(strings.ToUpper(method)),
				PathNameForm:   openapispec.NewNameForm(strings.ToUpper(path)),
				Description:    operation.Description,
				RouteNameSpace: data.RouteNameSpace,
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
					} else {
						requestSchemaName := strcase.SnakeCase(path) + "_" + strings.ToLower(method) + "_request"
						data.Schemas[requestSchemaName] = generateSchemaObject(requestSchemaName, requestSchema.Schema.Value)
						request.RequestSchemaName = generateName(requestSchemaName)
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
						request.Responses = append(request.Responses, &openapispec.Response{
							StatusCode: statusCode,
							Schema:     schema,
							Success:    success,
						})
					}
				}
			}
			data.Requests = append(data.Requests, &request)
		}
	}
}
