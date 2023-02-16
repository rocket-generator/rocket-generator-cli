package a_download_template

import (
	"github.com/fatih/color"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	url, err := getDownloadURL(payload.TemplateName)
	zipFilePath, err := downloadFile(*url, payload.ProjectBasePath)
	if err != nil {
		return nil, err
	}
	err = unzip(zipFilePath, payload.ProjectBasePath)
	if err != nil {
		return nil, err
	}
	_, zipFileName := filepath.Split(zipFilePath)
	slice := strings.Split(zipFileName, ".")
	err = os.Rename(slice[0], payload.ProjectName)
	if err != nil {
		return nil, err
	}

	err = os.Remove(zipFilePath)
	if err != nil {
		return nil, err
	}
	projectPath := path.Join(payload.ProjectBasePath, payload.ProjectName)
	payload.ProjectPath = projectPath

	typeMapperFilePath := filepath.Join(projectPath, "templates", "data", "types.json")
	typeMapper, err := data_mapper.Parse(typeMapperFilePath)
	if err == nil {
		red := color.New(color.FgGreen)
		_, _ = red.Println("Type mapping file found at: " + typeMapperFilePath)
		payload.TypeMapper = typeMapper
	} else {
		red := color.New(color.FgYellow)
		_, _ = red.Println("No type mapping file found at: " + typeMapperFilePath)
	}

	return payload, nil
}
