package handlers_test

import (
	"capital-gains/src/driven/operations"
	"testing"

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
	err := handler.Handle(command)

	// Then I expect no error
	assert.NoError(t, err)

	// And I expect the operation to be saved in the repository
	actual := repository.FindAll()

	assert.Len(t, actual, 1)

	// And I expect the operation to have the correct total value (100 * 10.00 = 1000.00)
	assert.Equal(t, 1000.00, actual[0].TotalValue().ToFloat64())
}
