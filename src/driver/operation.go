package driver

import (
	"capital-gains/src/application/commands"
	"strings"
)

const (
	buyOperationName  = "buy"
	sellOperationName = "sell"
)

type Operation struct {
	Quantity  int     `json:"quantity"`
	UnitCost  float64 `json:"unit-cost"`
	Operation string  `json:"operation"`
}

func (operation Operation) ToCommand() commands.Command {
	normalizedOperationName := strings.ToLower(strings.TrimSpace(operation.Operation))

	switch normalizedOperationName {
	case buyOperationName:
		return commands.NewRegisterBuy(operation.Quantity, operation.UnitCost)
	case sellOperationName:
		return commands.NewRegisterSell(operation.Quantity, operation.UnitCost)
	default:
		panic("unsupported operation")
	}
}
