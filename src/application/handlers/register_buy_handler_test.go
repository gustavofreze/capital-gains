package handlers_test

import (
	"testing"

	"capital-gains/src/driven/operations"

	"capital-gains/src/application/commands"
	"capital-gains/src/application/handlers"

	"github.com/stretchr/testify/assert"
)

func TestRegisterBuyHandlerGivenValidCommandWhenHandleThenBuyOperationIsPersisted(t *testing.T) {
	// Given that I have a command to register a buy operation
	command := commands.NewRegisterBuy(100, 10.00)

	// And I have a configured operations repository
	repository := operations.NewRepository()

	// When I handle the command with the buy handler
	handler := handlers.NewRegisterBuyHandler(repository)
	handler.Handle(command)

	// Then I expect the operation to be saved in the repository
	actual := repository.FindAll()

	assert.Len(t, actual, 1)
}
