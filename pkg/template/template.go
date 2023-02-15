package template

import (
	"bytes"
	"os"
	"path/filepath"
	textTemplate "text/template"
)

// GenerateFileFromTemplate ...
func GenerateFileFromTemplate(templateFilePath string, projectBasePath string, destinationFileDirectory string, data interface{}) error {

	_, file := filepath.Split(templateFilePath)

	file = file[:len(file)-len(filepath.Ext(file))]
	fileTemplate, err := textTemplate.New("filename_template").Parse(file)
	if err != nil {
		return err
	}
	fileNameBuffer := &bytes.Buffer{}
	err = fileTemplate.Execute(fileNameBuffer, data)
	if err != nil {
		return err
	}
	resultFileName := fileNameBuffer.String()

	contentTemplate, err := textTemplate.ParseFiles(templateFilePath)
	contentBuffer := &bytes.Buffer{}
	err = contentTemplate.Execute(contentBuffer, data)
	if err != nil {
		return err
	}
	err = os.WriteFile(destinationFileDirectory+string(os.PathSeparator)+resultFileName, contentBuffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}
