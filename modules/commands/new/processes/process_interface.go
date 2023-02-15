package processes

import (
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
)

type ProcessInterface interface {
	Execute(payload *newCommand.Payload) (*newCommand.Payload, error)
}
