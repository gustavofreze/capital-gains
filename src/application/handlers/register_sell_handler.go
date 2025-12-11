package handlers

import (
	"capital-gains/src/application/commands"
	"capital-gains/src/application/domain/models"
)

type RegisterSellHandler struct {
	position *models.Position
}

func NewRegisterSellHandler(position *models.Position) RegisterSellHandler {
	return RegisterSellHandler{
		position: position,
	}
}

func (handler RegisterSellHandler) Handle(command commands.RegisterSell) error {
	return nil
}
