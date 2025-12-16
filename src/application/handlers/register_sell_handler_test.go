package handlers_test

import (
	"capital-gains/src/driven/operations"
	"testing"

	"capital-gains/src/application/commands"
	"capital-gains/src/application/handlers"

	"github.com/stretchr/testify/assert"
)

func TestRegisterSellHandlerGivenValidCommandWhenHandleThenSellOperationIsPersisted(t *testing.T) {
	// Given that I have a command to register a sell operation
	command := commands.NewRegisterSell(100, 20.00)

	// And I have a configured operations repository
	repository := operations.NewRepository()

	// When I handle the command with the sell handler
	handler := handlers.NewRegisterSellHandler(repository)
	handler.Handle(command)

	// And I expect the operation to be saved in the repository
	actual := repository.FindAll()

	assert.Len(t, actual, 1)
}
