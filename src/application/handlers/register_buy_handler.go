package handlers

import (
	"capital-gains/src/application/commands"
	"capital-gains/src/application/domain/models"
	"capital-gains/src/application/ports/inbound"
	"capital-gains/src/application/ports/outbound"
)

var _ inbound.CommandHandler = (*RegisterBuyHandler)(nil)

type RegisterBuyHandler struct {
	operations outbound.Operations
}

func NewRegisterBuyHandler(operations outbound.Operations) RegisterBuyHandler {
	return RegisterBuyHandler{
		operations: operations,
	}
}

func (handler RegisterBuyHandler) Handle(command commands.Command) error {
	registerBuy := command.(commands.RegisterBuy)
	quantity := models.NewQuantity(registerBuy.Quantity())
	unitCost := models.NewMonetaryValue(registerBuy.UnitCost())

	buy := models.NewBuy(quantity, unitCost)

	handler.operations.Save(buy)
	return nil
}
