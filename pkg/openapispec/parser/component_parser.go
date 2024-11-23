package parser

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
	"github.com/stoewer/go-strcase"
)

func parseComponents(components openapi3.Components, api *objects.API, typeMapper *data_mapper.Mapper, orderedProperties *map[string]objects.OrderedProperties) {
	for name, schemaRef := range components.Schemas {
		specSchema := schemaRef.Value
		if specSchema == nil {
			continue
		}
		if specSchema.Type == nil || len(*specSchema.Type) == 0 || (*specSchema.Type)[0] != "object" {
			continue
		}
		schemaName := getSchemaNameFromSchema(name, schemaRef.Value)
		var targetOrderedProperties *objects.OrderedProperties
		if orderedProperties != nil {
			if properties, ok := (*orderedProperties)[schemaName]; ok {
				targetOrderedProperties = &properties
			}
		}
		api.Schemas[schemaName] = generateSchemaObject(schemaName, schemaRef.Value, typeMapper, targetOrderedProperties)
	}
}

func generateSchemaObject(schemaName string, schema *openapi3.Schema, typeMapper *data_mapper.Mapper, targetOrderedProperties *objects.OrderedProperties) *objects.Schema {

	schemaObject := objects.Schema{
		Name:        generateName(schemaName),
		Description: schema.Description,
	}
	requiredMap := map[string]bool{}
	for _, requiredColumn := range schema.Required {
		requiredMap[requiredColumn] = true
	}

	if targetOrderedProperties == nil {
		fmt.Println("No ordered properties for schema:", schemaName)
		targetOrderedProperties = &objects.OrderedProperties{
			Name:       schemaName,
			Properties: []string{},
		}
		for name, _ := range schema.Properties {
			targetOrderedProperties.Properties = append(targetOrderedProperties.Properties, name)
		}
	}
	for _, name := range targetOrderedProperties.Properties {
		property := schema.Properties[name]
		_, required := requiredMap[name]
		dataType := (*property.Value.Type)[0]
		switch dataType {
		case "array":
			itemName := getSchemaNameFromSchema(property.Value.Items.Ref, property.Value.Items.Value)
			item := property.Value.Items.Value
			schemaObject.Properties = append(schemaObject.Properties, &objects.Property{
				Name:          generateName(name),
				Type:          (*property.Value.Type)[0],
				ObjectType:    data_mapper.MapString(typeMapper, "database", dataType),
				CodeType:      data_mapper.MapString(typeMapper, "code", dataType),
				Description:   property.Value.Description,
				ArrayItemType: (*item.Type)[0],
				ArrayItemName: itemName,
				Required:      required,
			})
		case "object":
			propertyName := getSchemaNameFromSchema(property.Ref, property.Value)
			schemaObject.Properties = append(schemaObject.Properties, &objects.Property{
				Name:        generateName(name),
				Type:        (*property.Value.Type)[0],
				ObjectType:  data_mapper.MapString(typeMapper, "database", dataType),
				CodeType:    data_mapper.MapString(typeMapper, "code", dataType),
				Description: property.Value.Description,
				Reference:   propertyName,
				Required:    required,
			})
		default:
			schemaObject.Properties = append(schemaObject.Properties, &objects.Property{
				Name:        generateName(name),
				Type:        (*property.Value.Type)[0],
				ObjectType:  data_mapper.MapString(typeMapper, "database", dataType),
				CodeType:    data_mapper.MapString(typeMapper, "code", dataType),
				Description: property.Value.Description,
				Required:    required,
			})
		}
	}
	return &schemaObject
}

func getSchemaNameFromSchema(name string, schema *openapi3.Schema) string {
	elements := strings.Split(name, "/")
	schemaName := strcase.SnakeCase(elements[len(elements)-1])
	return schemaName
}
