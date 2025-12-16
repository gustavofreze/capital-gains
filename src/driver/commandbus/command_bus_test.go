package commandbus_test

import (
	"testing"

	"capital-gains/src/driver/commandbus"

	"capital-gains/test"

	"capital-gains/src/application/commands"

	"github.com/stretchr/testify/assert"
)

func TestCommandBusDispatchGivenRegisterBuyCommandWhenDispatchThenOnlyRegisterBuyHandlerIsInvoked(t *testing.T) {
	t.Parallel()

	// Given a command bus with all handlers
	registerBuyHandler := test.NewRegisterBuyHandlerMock()
	registerSellHandler := test.NewRegisterSellHandlerMock()
	calculateCapitalGainHandler := test.NewCalculateCapitalGainHandlerMock()
	commandBus := commandbus.NewCommandBus(
		registerBuyHandler,
		registerSellHandler,
		calculateCapitalGainHandler,
	)

	// And a register buy command
	registerBuyCommand := commands.NewRegisterBuy(100, 10.00)

	// When I dispatch the register buy command
	commandBus.Dispatch(registerBuyCommand)

	// Then I expect the register buy handler to be called once with the dispatched command
	assert.Equal(t, 1, registerBuyHandler.Calls())
	assert.Equal(t, registerBuyCommand, registerBuyHandler.FirstReceived())

	// And I expect the other handlers not to be called
	assert.Equal(t, 0, registerSellHandler.Calls())
	assert.Equal(t, 0, calculateCapitalGainHandler.Calls())
}

func TestCommandBusDispatchGivenRegisterSellCommandWhenDispatchThenOnlyRegisterSellHandlerIsInvoked(t *testing.T) {
	t.Parallel()

	// Given a command bus with all handlers
	registerBuyHandler := test.NewRegisterBuyHandlerMock()
	registerSellHandler := test.NewRegisterSellHandlerMock()
	calculateCapitalGainHandler := test.NewCalculateCapitalGainHandlerMock()
	commandBus := commandbus.NewCommandBus(
		registerBuyHandler,
		registerSellHandler,
		calculateCapitalGainHandler,
	)

	// And a register sell command
	registerSellCommand := commands.NewRegisterSell(100, 15.00)

	// When I dispatch the register sell command
	commandBus.Dispatch(registerSellCommand)

	// Then I expect the register sell handler to be called once with the dispatched command
	assert.Equal(t, 1, registerSellHandler.Calls())
	assert.Equal(t, registerSellCommand, registerSellHandler.FirstReceived())

	// And I expect the other handlers not to be called
	assert.Equal(t, 0, registerBuyHandler.Calls())
	assert.Equal(t, 0, calculateCapitalGainHandler.Calls())
}

func TestCommandBusDispatchGivenCalculateCapitalGainCommandWhenDispatchThenOnlyCalculateCapitalGainHandlerIsInvoked(t *testing.T) {
	t.Parallel()

	// Given a command bus with all handlers
	registerBuyHandler := test.NewRegisterBuyHandlerMock()
	registerSellHandler := test.NewRegisterSellHandlerMock()
	calculateCapitalGainHandler := test.NewCalculateCapitalGainHandlerMock()
	commandBus := commandbus.NewCommandBus(
		registerBuyHandler,
		registerSellHandler,
		calculateCapitalGainHandler,
	)

	// And a calculate capital gain command
	calculateCapitalGainCommand := commands.NewCalculateCapitalGain()

	// When I dispatch the calculate capital gain command
	commandBus.Dispatch(calculateCapitalGainCommand)

	// Then I expect the calculate capital gain handler to be called once with the dispatched command
	assert.Equal(t, 1, calculateCapitalGainHandler.Calls())
	assert.Equal(t, calculateCapitalGainCommand, calculateCapitalGainHandler.FirstReceived())

	// And I expect the other handlers not to be called
	assert.Equal(t, 0, registerBuyHandler.Calls())
	assert.Equal(t, 0, registerSellHandler.Calls())
}

func TestCommandBusDispatchGivenUnsupportedCommandWhenDispatchThenPanics(t *testing.T) {
	t.Parallel()

	// Given a command bus with all handlers
	registerBuyHandler := test.NewRegisterBuyHandlerMock()
	registerSellHandler := test.NewRegisterSellHandlerMock()
	calculateCapitalGainHandler := test.NewCalculateCapitalGainHandlerMock()
	commandBus := commandbus.NewCommandBus(
		registerBuyHandler,
		registerSellHandler,
		calculateCapitalGainHandler,
	)

	// And an unsupported command
	var unsupportedCommand commands.Command = nil

	// When I dispatch the unsupported command
	// Then I expect the application to panic
	assert.Panics(t, func() {
		commandBus.Dispatch(unsupportedCommand)
	})

	// And I expect no handler to be invoked
	assert.Equal(t, 0, registerBuyHandler.Calls())
	assert.Equal(t, 0, registerSellHandler.Calls())
	assert.Equal(t, 0, calculateCapitalGainHandler.Calls())
}
