package parser

import (
	"github.com/jinzhu/inflection"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec"
	"github.com/stoewer/go-strcase"
)

func generateName(name string) openapispec.Name {
	singular := inflection.Singular(name)
	plural := inflection.Plural(name)
	return openapispec.Name{
		Original: name,
		Default: openapispec.NameForm{
			Camel: strcase.LowerCamelCase(name),
			Title: strcase.UpperCamelCase(name),
			Snake: strcase.SnakeCase(name),
			Kebab: strcase.KebabCase(name),
		},
		Singular: openapispec.NameForm{
			Camel: strcase.LowerCamelCase(singular),
			Title: strcase.UpperCamelCase(singular),
			Snake: strcase.SnakeCase(singular),
			Kebab: strcase.KebabCase(singular),
		},
		Plural: openapispec.NameForm{
			Camel: strcase.LowerCamelCase(plural),
			Title: strcase.UpperCamelCase(plural),
			Snake: strcase.SnakeCase(plural),
			Kebab: strcase.KebabCase(plural),
		},
	}
}
