package service

import (
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
)

type Arguments struct {
	Type                      string
	Name                      string
	RelatedModelNames         []string
	RelatedModelWithCRUDNames []string
	RelatedResponse           *objects.Response
	IsAuthService             bool
	ProjectPath               string
	Debug                     bool
}
