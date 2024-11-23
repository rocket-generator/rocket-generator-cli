package parser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jinzhu/inflection"
	"github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
	"github.com/stoewer/go-strcase"
)

func findEntityIndex(name string, schema *objects.Schema) int {
	for index, entity := range schema.Entities {
		if entity.Name.Original == name {
			return index
		}
	}
	return -1
}

func findRelationColumnIndex(name string, table *objects.Entity) int {
	columnName := inflection.Singular(name)
	for index, column := range table.Columns {
		if column.Name.Original == columnName {
			return index
		}
	}
	return -1
}

func removeComment(content string) string {
	commentRegex := regexp.MustCompile(`(?ms)\/'.+?'\/`)
	return commentRegex.ReplaceAllString(content, "")
}

func checkAPIReturnable(column *objects.Column) bool {
	return true
}

func checkAPIUpdatable(column *objects.Column) bool {
	if column.Name.Original == "id" || column.Name.Original == "created_at" || column.Name.Original == "updated_at" || column.Name.Original == "deleted_at" {
		return false
	}
	return true
}

func getAPIType(column *objects.Column) string {
	if strings.HasPrefix(column.DataType.Original, "decimal") || strings.HasPrefix(column.DataType.Original, "numeric") {
		return "number"
	}
	switch column.DataType.Original {
	case "int":
		return "integer"
	case "bigint":
		return "string"
	case "text":
		return "string"
	case "boolean":
		return "boolean"
	case "float":
		return "number"
	case "jsonb":
		return "string"
	case "timestamp":
		return "integer"
	}
	return "string"
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

func convertDataTypeToDataTypeAndSize(dataType string) (string, int) {
	dataTypeRegex := regexp.MustCompile(`(?m)(\w+)\((\d+)\)`)
	foundDataType := dataTypeRegex.FindAllStringSubmatch(dataType, 1)
	if len(foundDataType) > 0 {
		size, err := strconv.Atoi(foundDataType[0][2])
		if err != nil {
			return dataType, 0
		}
		return foundDataType[0][1], size
	}
	return dataType, 0
}
