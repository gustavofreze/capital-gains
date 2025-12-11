package handlers

import (
	"capital-gains/src/application/commands"
	"capital-gains/src/application/domain/models"
)

type RegisterBuyHandler struct {
	position *models.Position
}

func NewRegisterBuyHandler(position *models.Position) RegisterBuyHandler {
	return RegisterBuyHandler{
		position: position,
	}
}

func (handler RegisterBuyHandler) Handle(command commands.RegisterBuy) error {
	return nil
}
