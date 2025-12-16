package handlers

import (
	"capital-gains/src/application/commands"
	"capital-gains/src/application/domain/models"
	"capital-gains/src/application/ports/outbound"
)

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

func (handler CalculateCapitalGainHandler) Handle(_ commands.CalculateCapitalGain) {
	operations := handler.operations.FindAll()

	capitalGain := models.NewCapitalGain()
	capitalGain.ApplyOperations(operations)

	handler.capitalGains.Save(capitalGain)
}
