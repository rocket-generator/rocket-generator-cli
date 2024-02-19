package parser

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
	"github.com/stoewer/go-strcase"
	"strings"
)

func parseComponents(components openapi3.Components, api *objects.API, typeMapper *data_mapper.Mapper) {
	for name, schemaRef := range components.Schemas {
		specSchema := schemaRef.Value
		if specSchema == nil {
			continue
		}
		if specSchema.Type != "object" {
			continue
		}
		schemaName := getSchemaNameFromSchema(name, schemaRef.Value)
		api.Schemas[schemaName] = generateSchemaObject(schemaName, schemaRef.Value, typeMapper)
	}
}

func generateSchemaObject(name string, schema *openapi3.Schema, typeMapper *data_mapper.Mapper) *objects.Schema {
	schemaObject := objects.Schema{
		Name:        generateName(name),
		Description: schema.Description,
	}
	requiredMap := map[string]bool{}
	for _, requiredColumn := range schema.Required {
		requiredMap[requiredColumn] = true
	}
	for name, property := range schema.Properties {
		_, required := requiredMap[name]
		switch property.Value.Type {
		case "array":
			itemName := getSchemaNameFromSchema(property.Value.Items.Ref, property.Value.Items.Value)
			item := property.Value.Items.Value
			schemaObject.Properties = append(schemaObject.Properties, &objects.Property{
				Name:          generateName(name),
				Type:          property.Value.Type,
				ObjectType:    data_mapper.MapString(typeMapper, "database", property.Value.Type),
				Description:   property.Value.Description,
				ArrayItemType: item.Type,
				ArrayItemName: itemName,
				Required:      required,
			})
		case "object":
			propertyName := getSchemaNameFromSchema(property.Ref, property.Value)
			schemaObject.Properties = append(schemaObject.Properties, &objects.Property{
				Name:        generateName(name),
				Type:        property.Value.Type,
				ObjectType:  data_mapper.MapString(typeMapper, "database", property.Value.Type),
				Description: property.Value.Description,
				Reference:   propertyName,
				Required:    required,
			})
		default:
			schemaObject.Properties = append(schemaObject.Properties, &objects.Property{
				Name:        generateName(name),
				Type:        property.Value.Type,
				ObjectType:  data_mapper.MapString(typeMapper, "database", property.Value.Type),
				Description: property.Value.Description,
				Required:    required,
			})
		}
	}
	return &schemaObject
}

func getSchemaNameFromSchema(name string, schema *openapi3.Schema) string {
	if schema.Title != "" {
		return schema.Title
	}
	elements := strings.Split(name, "/")
	schemaName := strcase.SnakeCase(elements[len(elements)-1])
	return schemaName
}
