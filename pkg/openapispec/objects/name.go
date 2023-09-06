package objects

import "github.com/stoewer/go-strcase"

// Name ...
type Name struct {
	Original string
	Default  NameForm
	Singular NameForm
	Plural   NameForm
}

// NameForm ...
type NameForm struct {
	Original string
	Camel    string
	Title    string
	Snake    string
	Kebab    string
}

func NewNameForm(name string) NameForm {
	return NameForm{
		Original: name,
		Camel:    strcase.LowerCamelCase(name),
		Title:    strcase.UpperCamelCase(name),
		Snake:    strcase.SnakeCase(name),
		Kebab:    strcase.KebabCase(name),
	}
}
