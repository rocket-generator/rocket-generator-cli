package g_build_admin_api

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	textTemplate "text/template"

	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
	"github.com/rocket-generator/rocket-generator-cli/pkg/error_handler"
	"github.com/rocket-generator/rocket-generator-cli/pkg/template"
)

func (process *Process) generateFileFromTemplate(entity objects.Entity, payload *newCommand.Payload) error {
	templatePath := filepath.Join(payload.ProjectPath, "templates", "admin_api")
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

			resultDirectoryRaw := filepath.Join(payload.ProjectPath, relativePath)
			// resultDirectory もテンプレートの処理を行う
			resultDirectory, err := processFileNameTemplate(resultDirectoryRaw, entity)
			if err != nil {
				error_handler.HandleError(err)
				return err
			}

			// ファイル名のテンプレート処理
			_, file := filepath.Split(path)
			file = file[:len(file)-len(filepath.Ext(file))]

			// テンプレートからファイル名を生成
			destinationFileName, err := processFileNameTemplate(file, entity)
			if err != nil {
				error_handler.HandleError(err)
				return err
			}
			if destinationFileName == "" {
				return nil
			}

			destinationPath := filepath.Join(resultDirectory, destinationFileName)
			// ファイルが既に存在する場合はスキップ
			if _, err := os.Stat(destinationPath); err == nil {
				fmt.Println("Skipping file:", destinationPath)
				return nil
			}

			// ファイルの親ディレクトリが存在しているかチェックし、なければ作成
			if _, err := os.Stat(resultDirectory); os.IsNotExist(err) {
				err := os.MkdirAll(resultDirectory, 0755)
				if err != nil {
					error_handler.HandleError(err)
					return err
				}
			}

			_, err = template.GenerateFileFromTemplate(path, payload.ProjectPath, resultDirectory, entity)
			if err != nil {
				error_handler.HandleError(err)
				return err
			}
		}
		return nil
	})
	return err
}

// processFileNameTemplate はテンプレートファイル名を処理して実際のファイル名を生成します
func processFileNameTemplate(templateFileName string, data interface{}) (string, error) {
	tmpl, err := textTemplate.New("filename_template").Parse(templateFileName)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	result := buf.String()
	if result == filepath.Ext(result) {
		return "", nil
	}

	return result, nil
}
