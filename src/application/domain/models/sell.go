package models

type Sell struct {
	unitCost UnitCost
	quantity Quantity
}

func NewSell(unitCost UnitCost, quantity Quantity) Sell {
	return Sell{
		unitCost: unitCost,
		quantity: quantity,
	}
}

func (sell Sell) Quantity() Quantity {
	return sell.quantity
}

func (sell Sell) TotalProceeds() MonetaryValue {
	return sell.unitCost.MultiplyBy(sell.quantity)
}

func (sell Sell) CalculateGrossCapitalGain(unitCost UnitCost) MonetaryValue {
	unitDifference := sell.unitCost.Subtract(unitCost)
	return unitDifference.MultiplyBy(sell.quantity)
}
