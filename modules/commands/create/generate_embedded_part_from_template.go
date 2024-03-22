package create

import (
	"github.com/rocket-generator/rocket-generator-cli/internal/utilities"
	"github.com/rocket-generator/rocket-generator-cli/pkg/error_handler"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
	"github.com/rocket-generator/rocket-generator-cli/pkg/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Entities struct {
	Requests []*objects.Request
}

func GenerateEmbeddedPartFromTemplate(projectPath string, targetType string, payload interface{}) error {
	templatePath := filepath.Join(projectPath, "templates", "create", targetType)
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
			targetFile := filepath.Join(projectPath, relativePath)
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
						replacement, err := template.GenerateStringFromTemplate(partialTemplateFullPath, payload)
						if err != nil {
							return err
						}
						originalContent, err := os.ReadFile(targetFile)
						if err != nil {
							return err
						}
						replacedText := utilities.GetFilenameWithoutExtension(partialTemplateFile.Name())
						placeHolder := "/* [" + replacedText + "] */"
						finalReplacement := *replacement + "\n" + placeHolder
						updatedText := strings.ReplaceAll(string(originalContent), placeHolder, finalReplacement)

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
