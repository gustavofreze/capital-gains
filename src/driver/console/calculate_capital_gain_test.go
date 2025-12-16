package console_test

import (
	"capital-gains/src/application/handlers"
	"capital-gains/src/driven/capitalgains"
	"capital-gains/src/driven/operations"
	"capital-gains/src/driver"
	"capital-gains/src/driver/console"
	"capital-gains/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCapitalGainPrintsExpectedTaxesForSingleInput(t *testing.T) {
	t.Parallel()

	// Given a sequence of market operations (buy/sell) to be processed
	payload := []map[string]any{
		{"operation": "buy", "unit-cost": 10.00, "quantity": 10000},
		{"operation": "sell", "unit-cost": 2.00, "quantity": 5000},
		{"operation": "sell", "unit-cost": 20.00, "quantity": 2000},
		{"operation": "sell", "unit-cost": 20.00, "quantity": 2000},
		{"operation": "sell", "unit-cost": 25.00, "quantity": 1000},
		{"operation": "buy", "unit-cost": 20.00, "quantity": 10000},
		{"operation": "sell", "unit-cost": 15.00, "quantity": 5000},
		{"operation": "sell", "unit-cost": 30.00, "quantity": 4350},
		{"operation": "sell", "unit-cost": 30.00, "quantity": 650},
	}
	defaultConsole := test.NewConsoleMock([]string{test.ToJson(payload)})

	// When processing these operations to calculate taxes
	operationsRepository := operations.NewRepository()
	capitalGainsRepository := capitalgains.NewRepository()
	calculateCapitalGains := console.NewCalculateCapitalGain(
		defaultConsole,
		handlers.NewRegisterBuyHandler(operationsRepository),
		handlers.NewRegisterSellHandler(operationsRepository),
		capitalGainsRepository,
		handlers.NewCalculateCapitalGainHandler(
			operationsRepository,
			capitalGainsRepository,
		),
	)
	calculateCapitalGains.Handle()

	// Then I expect the result to be written in a single output line
	assert.Len(t, defaultConsole.WrittenLines(), 1)

	// And I expect the tax values to match the specification output
	expectedTaxes := []driver.Tax{
		driver.NewTax(0.00),
		driver.NewTax(0.00),
		driver.NewTax(0.00),
		driver.NewTax(0.00),
		driver.NewTax(3000.00),
		driver.NewTax(0.00),
		driver.NewTax(0.00),
		driver.NewTax(3700.00),
		driver.NewTax(0.00),
	}
	assert.Equal(t, test.ToJson(expectedTaxes), defaultConsole.GetByIndex(0))
}

func TestCalculateCapitalGainPrintsExpectedTaxesForMultipleInputs(t *testing.T) {
	t.Parallel()

	// Given two input lines with sequences of market operations (buy/sell) to be processed
	firstOperations := []map[string]any{
		{"operation": "buy", "unit-cost": 10.00, "quantity": 100},
		{"operation": "sell", "unit-cost": 15.00, "quantity": 50},
		{"operation": "sell", "unit-cost": 15.00, "quantity": 50},
	}
	secondOperations := []map[string]any{
		{"operation": "buy", "unit-cost": 10.00, "quantity": 10000},
		{"operation": "sell", "unit-cost": 20.00, "quantity": 5000},
		{"operation": "sell", "unit-cost": 5.00, "quantity": 5000},
	}

	defaultConsole := test.NewConsoleMock([]string{
		test.ToJson(firstOperations),
		test.ToJson(secondOperations),
	})

	// When processing these operations to calculate taxes
	operationsRepository := operations.NewRepository()
	capitalGainsRepository := capitalgains.NewRepository()
	calculateCapitalGains := console.NewCalculateCapitalGain(
		defaultConsole,
		handlers.NewRegisterBuyHandler(operationsRepository),
		handlers.NewRegisterSellHandler(operationsRepository),
		capitalGainsRepository,
		handlers.NewCalculateCapitalGainHandler(
			operationsRepository,
			capitalGainsRepository,
		),
	)
	calculateCapitalGains.Handle()

	// Then I expect the result to be written in two output lines (one per input line)
	assert.Len(t, defaultConsole.WrittenLines(), 2)

	// And I expect the tax values of each input line to match the specification output
	firstExpected := []driver.Tax{
		driver.NewTax(0.00),
		driver.NewTax(0.00),
		driver.NewTax(0.00),
	}
	secondExpected := []driver.Tax{
		driver.NewTax(0.00),
		driver.NewTax(10000.00),
		driver.NewTax(0.00),
	}

	assert.Equal(t, test.ToJson(firstExpected), defaultConsole.GetByIndex(0))
	assert.Equal(t, test.ToJson(secondExpected), defaultConsole.GetByIndex(1))
}

func TestCalculateCapitalGainIgnoresTrailingEmptyLine(t *testing.T) {
	t.Parallel()

	// Given two input lines with sequences of market operations (buy/sell), plus a trailing empty line
	firstOperations := []map[string]any{
		{"operation": "buy", "unit-cost": 10.00, "quantity": 100},
		{"operation": "sell", "unit-cost": 15.00, "quantity": 50},
		{"operation": "sell", "unit-cost": 15.00, "quantity": 50},
	}
	secondOperations := []map[string]any{
		{"operation": "buy", "unit-cost": 10.00, "quantity": 10000},
		{"operation": "sell", "unit-cost": 20.00, "quantity": 5000},
		{"operation": "sell", "unit-cost": 5.00, "quantity": 5000},
	}

	defaultConsole := test.NewConsoleMock([]string{
		test.ToJson(firstOperations),
		test.ToJson(secondOperations),
		"",
	})

	// When processing these operations to calculate taxes
	operationsRepository := operations.NewRepository()
	capitalGainsRepository := capitalgains.NewRepository()
	calculateCapitalGains := console.NewCalculateCapitalGain(
		defaultConsole,
		handlers.NewRegisterBuyHandler(operationsRepository),
		handlers.NewRegisterSellHandler(operationsRepository),
		capitalGainsRepository,
		handlers.NewCalculateCapitalGainHandler(
			operationsRepository,
			capitalGainsRepository,
		),
	)
	calculateCapitalGains.Handle()

	// Then I expect the result to be written in two output lines (one per input line)
	assert.Len(t, defaultConsole.WrittenLines(), 2)

	// And I expect the tax values of each input line to match the specification output
	firstExpected := []driver.Tax{
		driver.NewTax(0.00),
		driver.NewTax(0.00),
		driver.NewTax(0.00),
	}
	secondExpected := []driver.Tax{
		driver.NewTax(0.00),
		driver.NewTax(10000.00),
		driver.NewTax(0.00),
	}

	assert.Equal(t, test.ToJson(firstExpected), defaultConsole.GetByIndex(0))
	assert.Equal(t, test.ToJson(secondExpected), defaultConsole.GetByIndex(1))
}

func TestCalculateCapitalGainPanicsWhenInputIsNotValidJson(t *testing.T) {
	t.Parallel()

	// Given an input line that cannot be parsed as a JSON list of operations
	defaultConsole := test.NewConsoleMock([]string{"this is not json"})

	// When handling the input
	operationsRepository := operations.NewRepository()
	capitalGainsRepository := capitalgains.NewRepository()
	calculateCapitalGains := console.NewCalculateCapitalGain(
		defaultConsole,
		handlers.NewRegisterBuyHandler(operationsRepository),
		handlers.NewRegisterSellHandler(operationsRepository),
		capitalGainsRepository,
		handlers.NewCalculateCapitalGainHandler(
			operationsRepository,
			capitalGainsRepository,
		),
	)

	// Then I expect the application to panic
	assert.Panics(t, func() {
		calculateCapitalGains.Handle()
	})
}
