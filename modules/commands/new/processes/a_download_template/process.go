package a_download_template

import (
	"errors"
	"github.com/fatih/color"
	cp "github.com/otiai10/copy"
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
	if err != nil {
		return nil, err
	}
	projectPath := path.Join(payload.ProjectBasePath, payload.ProjectName)
	if _, err := os.Stat(projectPath); err == nil {
		return nil, errors.New("project directory already exists:" + projectPath)
	}

	if strings.HasPrefix(*url, "http") {
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
	} else if _, err := os.Stat(*url); err == nil {
		err = cp.Copy(*url, projectPath)
		if err != nil {
			return nil, err
		}
		gitDirectoryPath := filepath.Join(projectPath, ".git")
		// delete .git directory to disconnect from the template
		if _, err := os.Stat(gitDirectoryPath); err == nil {
			err = os.RemoveAll(gitDirectoryPath)
			if err != nil {
				return nil, err
			}
		}
	}

	payload.ProjectPath = projectPath

	typeMapperFilePath := filepath.Join(projectPath, "templates", "data", "types.json")
	typeMapper, err := data_mapper.Parse(typeMapperFilePath)
	if err == nil {
		green := color.New(color.FgGreen)
		_, _ = green.Println("Type mapping file found at: " + typeMapperFilePath)
		payload.TypeMapper = typeMapper
	} else {
		yellow := color.New(color.FgYellow)
		_, _ = yellow.Println("No type mapping file found at: " + typeMapperFilePath)
	}

	return payload, nil
}
