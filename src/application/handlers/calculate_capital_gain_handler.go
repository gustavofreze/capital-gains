package handlers

import (
	"capital-gains/src/application/commands"
	"capital-gains/src/application/domain/models"
	"capital-gains/src/application/ports/inbound"
	"capital-gains/src/application/ports/outbound"
)

var _ inbound.CommandHandler = (*CalculateCapitalGainHandler)(nil)

type CalculateCapitalGainHandler struct {
	operations   outbound.Operations
	capitalGains outbound.CapitalGains
}

func NewCalculateCapitalGainHandler(
	operations outbound.Operations,
	capitalGains outbound.CapitalGains,
) CalculateCapitalGainHandler {
	return CalculateCapitalGainHandler{
		operations:   operations,
		capitalGains: capitalGains,
	}
}

func (handler CalculateCapitalGainHandler) Handle(command commands.Command) error {
	_ = command.(commands.CalculateCapitalGain)
	operations := handler.operations.FindAll()

	capitalGain := models.NewCapitalGains()
	capitalGain.ApplyOperations(operations)

	handler.capitalGains.Save(capitalGain)
	return nil
}
