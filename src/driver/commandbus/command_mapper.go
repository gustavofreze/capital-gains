package commandbus

import (
	"capital-gains/src/application/commands"
	"capital-gains/src/driver"
)

type CommandMapper struct {
	request driver.Request
}

func NewCommandMapper(request driver.Request) *CommandMapper {
	return &CommandMapper{request: request}
}

func (mapper *CommandMapper) Map() []commands.Command {
	operations := mapper.request.Operations()
	commandsToHandle := make([]commands.Command, len(operations))

	for index, operation := range operations {
		commandsToHandle[index] = operation.ToCommand()
	}

	return commandsToHandle
}
