package console

import (
	"capital-gains/src/application/commands"
	"capital-gains/src/application/ports/inbound"
	"capital-gains/src/application/ports/outbound"
	"capital-gains/src/driver"
	"capital-gains/src/driver/commandbus"
)

type CalculateCapitalGain struct {
	capitalGains      outbound.CapitalGains
	commandBus        *commandbus.CommandBus
	operationsConsole *OperationsConsole
}

func NewCalculateCapitalGain(
	console Console,
	registerBuy inbound.RegisterBuy,
	registerSell inbound.RegisterSell,
	capitalGains outbound.CapitalGains,
	calculateCapitalGain inbound.CalculateCapitalGain,
) *CalculateCapitalGain {
	commandBus := commandbus.NewCommandBus(registerBuy, registerSell, calculateCapitalGain)
	operationsConsole := NewOperationsConsole(console)

	return &CalculateCapitalGain{
		capitalGains:      capitalGains,
		commandBus:        commandBus,
		operationsConsole: operationsConsole,
	}
}

func (calculateCapitalGain *CalculateCapitalGain) Handle() {
	requests := calculateCapitalGain.operationsConsole.ReadRequests()

	for _, request := range requests {
		commandFactory := commandbus.NewCommandMapper(request)
		commandsToHandle := commandFactory.Map()

		for _, command := range commandsToHandle {
			calculateCapitalGain.commandBus.Dispatch(command)
		}

		calculateCapitalGain.commandBus.Dispatch(commands.NewCalculateCapitalGain())

		taxes := calculateCapitalGain.capitalGains.FindAll()
		response := driver.NewResponse(taxes)

		calculateCapitalGain.operationsConsole.WriteResponse(response)
	}
}
