package commands

var _ Command = (*RegisterSell)(nil)

type RegisterSell struct {
	quantity int
	unitCost float64
}

func NewRegisterSell(quantity int, unitCost float64) RegisterSell {
	return RegisterSell{
		quantity: quantity,
		unitCost: unitCost,
	}
}

func (command RegisterSell) Quantity() int {
	return command.quantity
}

func (command RegisterSell) UnitCost() float64 {
	return command.unitCost
}
