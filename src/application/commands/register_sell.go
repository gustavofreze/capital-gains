package commands

type RegisterSell struct {
	unitCost float64
	quantity int64
}

func NewRegisterSell(unitCost float64, quantity int64) RegisterSell {
	return RegisterSell{
		unitCost: unitCost,
		quantity: quantity,
	}
}

func (command RegisterSell) UnitCost() float64 {
	return command.unitCost
}

func (command RegisterSell) Quantity() int64 {
	return command.quantity
}
