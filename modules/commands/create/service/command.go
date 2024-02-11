package service

import "github.com/stoewer/go-strcase"

type Command struct {
}

func (c *Command) Execute(arguments Arguments) error {
	payload := Payload{
		Name:          generateName(removePostfix(arguments.ServiceName)),
		RelatedModels: []Name{},
		ProjectPath:   arguments.ProjectPath,
		Debug:         arguments.Debug,
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

func removePostfix(s string) string {
	lower := strcase.LowerCamelCase(s)
	if len(lower) > 7 && lower[len(lower)-7:] == "Service" {
		return s[:len(s)-3]
	}
	return s
}
