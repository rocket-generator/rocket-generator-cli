package parser

import (
	"github.com/jinzhu/inflection"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
	"github.com/stoewer/go-strcase"
)

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
