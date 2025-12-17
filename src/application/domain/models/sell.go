package models

type Sell struct {
	quantity Quantity
	unitCost MonetaryValue
}

func NewSell(quantity Quantity, unitCost MonetaryValue) Sell {
	return Sell{
		quantity: quantity,
		unitCost: unitCost,
	}
}

func (sell Sell) ApplyTo(position *Position) Tax {
	return position.Sell(sell.quantity, sell.unitCost)
}
