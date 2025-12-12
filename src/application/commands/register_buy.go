package commands

var _ Command = (*RegisterBuy)(nil)

type RegisterBuy struct {
	quantity int
	unitCost float64
}

func NewRegisterBuy(quantity int, unitCost float64) RegisterBuy {
	return RegisterBuy{
		quantity: quantity,
		unitCost: unitCost,
	}
}

func (command RegisterBuy) Quantity() int {
	return command.quantity
}

func (command RegisterBuy) UnitCost() float64 {
	return command.unitCost
}
