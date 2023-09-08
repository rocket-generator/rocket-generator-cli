package d_build_app_api

import (
	"fmt"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/pkg/error_handler"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
	"github.com/rocket-generator/rocket-generator-cli/pkg/template"
	"io/fs"
	"os"
	"path/filepath"
)

func (process *Process) generateFileFromTemplate(request objects.Request, payload *newCommand.Payload) error {
	templatePath := filepath.Join(payload.ProjectPath, "templates", "app_api")
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
			resultPath, err := template.GenerateFileFromTemplate(path, payload.ProjectPath, resultDirectory, request)
			if err != nil {
				error_handler.HandleError(err)
				return err
			}
			if resultPath != nil {
				fmt.Println("Generated file: ", *resultPath)
			}
		}
		return nil
	})
	return err
}
