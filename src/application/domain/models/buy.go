package models

type Buy struct {
	quantity Quantity
	unitCost MonetaryValue
}

func NewBuy(quantity Quantity, unitCost MonetaryValue) Buy {
	return Buy{
		quantity: quantity,
		unitCost: unitCost,
	}
}

func (buy Buy) ApplyTo(position *Position) Tax {
	position.Buy(buy.quantity, buy.unitCost)

	return NewTax(NewZeroMonetaryValue())
}
