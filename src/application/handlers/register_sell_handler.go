package handlers

import (
	"capital-gains/src/application/commands"
	"capital-gains/src/application/domain/models"
	"capital-gains/src/application/ports/outbound"
)

type RegisterSellHandler struct {
	operations outbound.Operations
}

func NewRegisterSellHandler(operations outbound.Operations) RegisterSellHandler {
	return RegisterSellHandler{
		operations: operations,
	}
}

func (handler RegisterSellHandler) Handle(command commands.RegisterSell) {
	quantity := models.NewQuantity(command.Quantity())
	unitCost := models.NewMonetaryValue(command.UnitCost())

	sell := models.NewSell(quantity, unitCost)

	handler.operations.Save(sell)
}
