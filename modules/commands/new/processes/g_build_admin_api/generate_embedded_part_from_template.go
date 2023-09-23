package g_build_admin_api

import (
	"github.com/rocket-generator/rocket-generator-cli/internal/utilities"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
	"github.com/rocket-generator/rocket-generator-cli/pkg/error_handler"
	"github.com/rocket-generator/rocket-generator-cli/pkg/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Entities struct {
	Entities []*objects.Entity
}

func (process *Process) generateEmbeddedPartFromTemplate(entities []*objects.Entity, payload *newCommand.Payload) error {
	templatePath := filepath.Join(payload.ProjectPath, "templates", "admin_api")
	if _, err := os.Stat(templatePath); err != nil {
		return err
	}
	err := filepath.Walk(templatePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		extension := filepath.Ext(path)

		if extension == ".embed" && info.IsDir() {
			_, file := filepath.Split(path)
			targetPath := path[:len(path)-len(filepath.Ext(file))]
			relativePath, err := filepath.Rel(templatePath, targetPath)
			if err != nil {
				return err
			}
			targetFile := filepath.Join(payload.ProjectPath, relativePath)
			if _, err := os.Stat(targetFile); err == nil {
				files, err := os.ReadDir(path)
				if err != nil {
					return err
				}
				for _, partialTemplateFile := range files {
					if partialTemplateFile.IsDir() {
						continue
					}
					ptmplExtension := filepath.Ext(partialTemplateFile.Name())
					if ptmplExtension == ".ptmpl" && !partialTemplateFile.IsDir() {
						partialTemplateFullPath := filepath.Join(path, partialTemplateFile.Name())
						replacement, err := template.GenerateStringFromTemplate(partialTemplateFullPath, Entities{
							Entities: entities,
						})
						if err != nil {
							return err
						}
						originalContent, err := os.ReadFile(targetFile)
						if err != nil {
							return err
						}
						replacedText := utilities.GetFilenameWithoutExtension(partialTemplateFile.Name())
						placeHolder := "/* [" + replacedText + "] */"
						updatedText := strings.ReplaceAll(string(originalContent), placeHolder, *replacement)

						err = os.WriteFile(targetFile, []byte(updatedText), os.ModePerm)
						if err != nil {
							return err
						}
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		error_handler.HandleError(err)
		return err
	}
	return nil
}
