package api

import (
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
)

type Arguments struct {
	Type                 string
	Path                 string
	Method               string
	ApiFileName          string
	ApiInfoFileName      string
	ServiceFileName      string
	DatabaseFileName     string
	DatabaseInfoFileName string
	Request              *objects.Request
	ApiSpec              *objects.API
	TypeMapper           *data_mapper.Mapper
	ProjectPath          string
	Debug                bool
}
