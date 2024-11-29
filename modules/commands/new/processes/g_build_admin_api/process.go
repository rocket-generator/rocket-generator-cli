package g_build_admin_api

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	for _, entity := range payload.DatabaseSchema.Entities {
		ignoreKey := strings.ToLower(entity.Name.Original)
		// Ignore if ignoreKey is in payload.IgnoreList
		fmt.Println("Check ignoreKey: " + ignoreKey)
		if payload.IgnoreList != nil {
			for key := range payload.IgnoreList.Tables {
				fmt.Println("Ignore Entry: " + key)
			}
			if _, ok := payload.IgnoreList.Tables[ignoreKey]; ok {
				yellow := color.New(color.FgYellow)
				_, _ = yellow.Println("* Ignore db table: " + ignoreKey)
				continue
			}
		}
		if err := process.generateFileFromTemplate(*entity, payload); err != nil {
			return nil, err
		}
		err := process.generateEmbeddedPartFromTemplate(*entity, payload)
		if err != nil {
			return nil, err
		}
	}
	return payload, nil
}
