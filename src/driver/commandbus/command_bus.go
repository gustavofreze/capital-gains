package commandbus

import (
	"capital-gains/src/application/commands"
	"capital-gains/src/application/ports/inbound"
)

type CommandBus struct {
	registerBuy          inbound.RegisterBuy
	registerSell         inbound.RegisterSell
	calculateCapitalGain inbound.CalculateCapitalGain
}

func NewCommandBus(
	registerBuy inbound.RegisterBuy,
	registerSell inbound.RegisterSell,
	calculateCapitalGain inbound.CalculateCapitalGain,
) *CommandBus {
	return &CommandBus{
		registerBuy:          registerBuy,
		registerSell:         registerSell,
		calculateCapitalGain: calculateCapitalGain,
	}
}

func (commandBus *CommandBus) Dispatch(command commands.Command) {
	switch typedCommand := command.(type) {
	case commands.RegisterBuy:
		commandBus.registerBuy.Handle(typedCommand)
	case commands.RegisterSell:
		commandBus.registerSell.Handle(typedCommand)
	case commands.CalculateCapitalGain:
		commandBus.calculateCapitalGain.Handle(typedCommand)
	default:
		panic("unsupported command")
	}
}
