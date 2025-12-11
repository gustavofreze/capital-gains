package commands

type RegisterBuy struct {
	unitCost float64
	quantity int64
}

func NewRegisterBuy(unitCost float64, quantity int64) RegisterBuy {
	return RegisterBuy{
		unitCost: unitCost,
		quantity: quantity,
	}
}

func (command RegisterBuy) UnitCost() float64 {
	return command.unitCost
}

func (command RegisterBuy) Quantity() int64 {
	return command.quantity
}
