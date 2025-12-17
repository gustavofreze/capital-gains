package handlers_test

import (
	"testing"

	"capital-gains/test"

	"capital-gains/src/application/commands"
	"capital-gains/src/application/handlers"
	"capital-gains/src/driven/capitalgains"
	"capital-gains/src/driven/operations"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCapitalGainHandlerGivenRegisteredOperationsWhenHandleThenCapitalGainIsPersistedWithExpectedTaxes(t *testing.T) {
	// Given that I have a configured operations repository
	operationsRepository := operations.NewRepository()

	// And I have a handler to register buy operations
	registerBuyHandler := handlers.NewRegisterBuyHandler(operationsRepository)

	// And I have a handler to register sell operations
	registerSellHandler := handlers.NewRegisterSellHandler(operationsRepository)

	// And I register a buy operation of 10000 units at 10.00
	buyCommand := commands.NewRegisterBuy(10000, 10.00)
	registerBuyHandler.Handle(buyCommand)

	// And I register a sell operation of 5000 units at 20.00
	sellCommand := commands.NewRegisterSell(5000, 20.00)
	registerSellHandler.Handle(sellCommand)

	// And I have a configured capital gain repository
	capitalGainRepository := capitalgains.NewRepository()

	// And I have a handler to calculate the capital gain using the registered operations
	calculateHandler := handlers.NewCalculateCapitalGainHandler(operationsRepository, capitalGainRepository)

	// And I have a command to calculate the capital gain
	calculateCommand := commands.NewCalculateCapitalGain()

	// When I handle the calculate capital gain command
	calculateHandler.Handle(calculateCommand)

	// Then I expect one capital gain result to be stored in the capital gain repository
	capitalGains := capitalGainRepository.FindAll()

	assert.Len(t, capitalGains, 1)

	// And I expect the stored capital gain to have the correct tax events:
	//   - zero tax for the buy operation
	//   - 10000.00 tax for the sell operation
	taxEvents := capitalGains[0].Events()
	taxAmounts := test.TaxAmountsFromEvents(taxEvents)

	expectedTaxAmounts := []float64{
		0.00,     // buy
		10000.00, // sell
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}
