package service

import (
	"github.com/rocket-generator/rocket-generator-cli/pkg/error_handler"
	"github.com/rocket-generator/rocket-generator-cli/pkg/template"
	"io/fs"
	"os"
	"path/filepath"
)

func (c *Command) generateFileFromTemplate(payload *Payload) error {
	templatePath := filepath.Join(payload.ProjectPath, "templates", "create", "service")
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
			_, err = template.GenerateFileFromTemplate(path, payload.ProjectPath, resultDirectory, payload)
			if err != nil {
				error_handler.HandleError(err)
				return err
			}
		}
		return nil
	})
	return err
}
