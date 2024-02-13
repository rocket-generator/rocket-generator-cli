package create

import (
	"github.com/stoewer/go-strcase"
	"strings"
)

type SubCommand struct {
}

func (c *SubCommand) Execute(arguments Arguments) error {
	payload := Payload{
		Type:            arguments.Type,
		Name:            generateName(removePostfix(arguments.Name, strcase.UpperCamelCase(arguments.Type))),
		RelatedModels:   []Name{},
		RelatedResponse: arguments.RelatedResponse,
		ProjectPath:     arguments.ProjectPath,
		Debug:           arguments.Debug,
	}

	for _, relatedModelName := range arguments.RelatedModelNames {
		payload.RelatedModels = append(payload.RelatedModels, generateName(relatedModelName))
	}

	err := c.generateFileFromTemplate(&payload)
	if err != nil {
		return err
	}
	err = c.generateEmbeddedPartFromTemplate(&payload)
	if err != nil {
		return err
	}
	return nil
}

func removePostfix(s string, postfix string) string {
	lower := strcase.LowerCamelCase(s)
	if strings.HasSuffix(lower, postfix) {
		return strings.TrimSuffix(s, postfix)
	}
	return s
}
