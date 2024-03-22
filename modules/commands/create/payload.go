package create

import (
	"github.com/jinzhu/inflection"
	"github.com/stoewer/go-strcase"
)

// Name ...
type Name struct {
	Original string
	Default  NameForm
	Singular NameForm
	Plural   NameForm
}

// NameForm ...
type NameForm struct {
	Camel string
	Title string
	Snake string
	Kebab string
}

func GenerateName(name string) Name {
	singular := inflection.Singular(name)
	plural := inflection.Plural(name)
	return Name{
		Original: name,
		Default: NameForm{
			Camel: strcase.LowerCamelCase(name),
			Title: strcase.UpperCamelCase(name),
			Snake: strcase.SnakeCase(name),
			Kebab: strcase.KebabCase(name),
		},
		Singular: NameForm{
			Camel: strcase.LowerCamelCase(singular),
			Title: strcase.UpperCamelCase(singular),
			Snake: strcase.SnakeCase(singular),
			Kebab: strcase.KebabCase(singular),
		},
		Plural: NameForm{
			Camel: strcase.LowerCamelCase(plural),
			Title: strcase.UpperCamelCase(plural),
			Snake: strcase.SnakeCase(plural),
			Kebab: strcase.KebabCase(plural),
		},
	}
}
