package create

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	textTemplate "text/template"

	"github.com/rocket-generator/rocket-generator-cli/pkg/error_handler"
	"github.com/rocket-generator/rocket-generator-cli/pkg/template"
)

func GenerateFileFromTemplate(projectPath string, targetType string, payload interface{}) error {
	templatePath := filepath.Join(projectPath, "templates", "create", targetType)
	if _, err := os.Stat(templatePath); err != nil {
		// No template directory. Just skip
		fmt.Println("[SKIP] No template directory for ", targetType)
		return nil
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
			resultDirectoryRaw := filepath.Join(projectPath, relativePath)
			// resultDirectory もテンプレートの処理を行う
			resultDirectory, err := processFileNameTemplate(resultDirectoryRaw, payload)
			if err != nil {
				error_handler.HandleError(err)
				return err
			}

			// ファイル名のテンプレート処理
			_, file := filepath.Split(path)
			file = file[:len(file)-len(filepath.Ext(file))]

			// テンプレートからファイル名を生成
			destinationFileName, err := processFileNameTemplate(file, payload)
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
				fmt.Println("Creating directory:", resultDirectory)
				err := os.MkdirAll(resultDirectory, 0755)
				if err != nil {
					error_handler.HandleError(err)
					return err
				}
			}

			_, err = template.GenerateFileFromTemplate(path, projectPath, resultDirectory, payload)
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
