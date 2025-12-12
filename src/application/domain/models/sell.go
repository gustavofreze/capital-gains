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

func (sell Sell) Quantity() Quantity {
	return sell.quantity
}

func (sell Sell) TotalValue() MonetaryValue {
	return sell.unitCost.MultiplyBy(sell.quantity.ToFloat())
}

func (sell Sell) CalculateGrossCapitalGain(averageUnitCost MonetaryValue) MonetaryValue {
	value := sell.unitCost.Subtract(averageUnitCost)
	product := value.MultiplyBy(sell.quantity.ToFloat())

	return product
}
