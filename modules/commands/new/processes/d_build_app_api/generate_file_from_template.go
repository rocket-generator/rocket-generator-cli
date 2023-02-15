package d_build_app_api

import (
	"fmt"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec"
	"github.com/rocket-generator/rocket-generator-cli/pkg/template"
	"io/fs"
	"os"
	"path/filepath"
)

func (process *Process) generateFileFromTemplate(request openapispec.Request, payload *newCommand.Payload) error {
	templatePath := filepath.Join(payload.ProjectPath, "templates", "app_api")
	if _, err := os.Stat(templatePath); err != nil {
		return err
	}
	err := filepath.Walk(templatePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		extension := filepath.Ext(path)
		if extension == ".tmpl" && !info.IsDir() {
			relativePath, err := filepath.Rel(templatePath, filepath.Dir(path))
			if err != nil {
				return err
			}
			resultDirectory := filepath.Join(payload.ProjectPath, relativePath)
			err = template.GenerateFileFromTemplate(path, payload.ProjectPath, resultDirectory, request)
			if err != nil {
				return err
			}
			fmt.Println("Generated file: ", resultDirectory)
		}
		return nil
	})
	return err
}
