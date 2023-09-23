package f_build_database

import (
	"fmt"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
	"github.com/rocket-generator/rocket-generator-cli/pkg/error_handler"
	"github.com/rocket-generator/rocket-generator-cli/pkg/template"
	"io/fs"
	"os"
	"path/filepath"
)

func (process *Process) generateFileFromTemplate(entity objects.Entity, payload *newCommand.Payload) error {
	templatePath := filepath.Join(payload.ProjectPath, "templates", "database")
	if _, err := os.Stat(templatePath); err != nil {
		return err
	}
	err := filepath.Walk(templatePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			error_handler.HandleError(err)
			return err
		}
		extension := filepath.Ext(path)
		if extension == ".tmpl" && !info.IsDir() {
			relativePath, err := filepath.Rel(templatePath, filepath.Dir(path))
			if err != nil {
				error_handler.HandleError(err)
				return err
			}
			resultDirectory := filepath.Join(payload.ProjectPath, relativePath)
			_, err = template.GenerateFileFromTemplate(path, payload.ProjectPath, resultDirectory, entity)
			if err != nil {
				error_handler.HandleError(err)
				return err
			}
		}
		return nil
	})
	fmt.Println("")
	return err
}
