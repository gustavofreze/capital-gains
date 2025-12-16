package handlers

import (
	"capital-gains/src/application/commands"
	"capital-gains/src/application/domain/models"
	"capital-gains/src/application/ports/outbound"
)

type RegisterBuyHandler struct {
	operations outbound.Operations
}

func NewRegisterBuyHandler(operations outbound.Operations) RegisterBuyHandler {
	return RegisterBuyHandler{
		operations: operations,
	}
}

func (handler RegisterBuyHandler) Handle(command commands.RegisterBuy) {
	quantity := models.NewQuantity(command.Quantity())
	unitCost := models.NewMonetaryValue(command.UnitCost())

	buy := models.NewBuy(quantity, unitCost)

	handler.operations.Save(buy)
}
