package handlers

import (
	"capital-gains/src/application/commands"
	"capital-gains/src/application/domain/models"
	"capital-gains/src/application/ports/inbound"
	"capital-gains/src/application/ports/outbound"
)

var _ inbound.CommandHandler = (*RegisterSellHandler)(nil)

type RegisterSellHandler struct {
	operations outbound.Operations
}

func NewRegisterSellHandler(operations outbound.Operations) RegisterSellHandler {
	return RegisterSellHandler{
		operations: operations,
	}
}

func (handler RegisterSellHandler) Handle(command commands.Command) error {
	registerSell := command.(commands.RegisterSell)
	quantity := models.NewQuantity(registerSell.Quantity())
	unitCost := models.NewMonetaryValue(registerSell.UnitCost())

	sell := models.NewSell(quantity, unitCost)

	handler.operations.Save(sell)
	return nil
}
