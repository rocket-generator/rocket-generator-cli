package model

import (
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	databaseObject "github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
)

type Arguments struct {
	Type             string
	Name             string
	DatabaseFileName string
	ProjectPath      string
	DatabaseSchema   *databaseObject.Schema
	TypeMapper       *data_mapper.Mapper
	Debug            bool
}
