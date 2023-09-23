package template

import (
	"bytes"
	"fmt"
	"github.com/rocket-generator/rocket-generator-cli/internal/utilities"
	"github.com/rocket-generator/rocket-generator-cli/pkg/error_handler"
	"os"
	"path/filepath"
	textTemplate "text/template"
)

// GenerateFileFromTemplate ...
func GenerateFileFromTemplate(templateFilePath string, projectBasePath string, destinationFileDirectory string, data interface{}) (*string, error) {

	// Convert File Name
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
	if resultFileName == filepath.Ext(resultFileName) {
		return nil, nil
	}

	// Convert Project Path
	pathTemplate, err := textTemplate.New("path_template").Parse(destinationFileDirectory)
	if err != nil {
		fmt.Println(file)
		error_handler.HandleError(err)
		return nil, err
	}
	filePathBuffer := &bytes.Buffer{}
	err = pathTemplate.Execute(filePathBuffer, data)
	if err != nil {
		error_handler.HandleError(err)
		return nil, err
	}
	resultFilePath := filePathBuffer.String()

	err = utilities.CreateDirectoryIfNotExists(resultFilePath)
	if err != nil {
		error_handler.HandleError(err)
		return nil, err
	}

	// Generate Content
	contentTemplate, err := textTemplate.ParseFiles(templateFilePath)
	contentBuffer := &bytes.Buffer{}
	err = contentTemplate.Execute(contentBuffer, data)
	if err != nil {
		error_handler.HandleError(err)
		return nil, err
	}
	err = os.WriteFile(resultFilePath+string(os.PathSeparator)+resultFileName, contentBuffer.Bytes(), 0644)
	if err != nil {
		error_handler.HandleError(err)
		return nil, err
	}
	return &resultFileName, nil
}

func GenerateStringFromTemplate(templateFilePath string, data interface{}) (*string, error) {
	contentTemplate, err := textTemplate.ParseFiles(templateFilePath)
	if err != nil {
		return nil, err
	}
	contentBuffer := &bytes.Buffer{}
	err = contentTemplate.Execute(contentBuffer, data)
	if err != nil {
		error_handler.HandleError(err)
		return nil, err
	}
	result := contentBuffer.String()
	return &result, nil
}
