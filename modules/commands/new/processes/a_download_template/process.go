package a_download_template

import (
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
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
	return payload, nil
}
