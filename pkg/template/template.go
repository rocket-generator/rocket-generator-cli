package template

import (
	"bytes"
	"fmt"
	"github.com/rocket-generator/rocket-generator-cli/pkg/error_handler"
	"os"
	"path/filepath"
	textTemplate "text/template"
)

// GenerateFileFromTemplate ...
func GenerateFileFromTemplate(templateFilePath string, projectBasePath string, destinationFileDirectory string, data interface{}) (*string, error) {
	fmt.Println("Generating file from template: " + templateFilePath)
	_, file := filepath.Split(templateFilePath)

	file = file[:len(file)-len(filepath.Ext(file))]
	fileTemplate, err := textTemplate.New("filename_template").Parse(file)
	if err != nil {
		fmt.Println(file)
		error_handler.HandleError(err)
		return nil, err
	}
	fileNameBuffer := &bytes.Buffer{}
	err = fileTemplate.Execute(fileNameBuffer, data)
	if err != nil {
		error_handler.HandleError(err)
		return nil, err
	}
	resultFileName := fileNameBuffer.String()

	contentTemplate, err := textTemplate.ParseFiles(templateFilePath)
	contentBuffer := &bytes.Buffer{}
	fmt.Println("Generating file from template: " + resultFileName)
	err = contentTemplate.Execute(contentBuffer, data)
	if err != nil {
		error_handler.HandleError(err)
		return nil, err
	}
	err = os.WriteFile(destinationFileDirectory+string(os.PathSeparator)+resultFileName, contentBuffer.Bytes(), 0644)
	if err != nil {
		error_handler.HandleError(err)
		return nil, err
	}
	return &resultFileName, nil
}
