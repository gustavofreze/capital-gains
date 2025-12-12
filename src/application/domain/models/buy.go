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

func (buy Buy) Quantity() Quantity {
	return buy.quantity
}

func (buy Buy) TotalValue() MonetaryValue {
	return buy.unitCost.MultiplyBy(buy.quantity.ToFloat())
}

func (buy Buy) CalculateWeightedAverageUnitCost(quantity Quantity, unitCost MonetaryValue) MonetaryValue {
	totalCost := unitCost.MultiplyBy(quantity.ToFloat())
	combinedQuantity := quantity.Add(buy.quantity)
	combinedTotalCost := totalCost.Add(buy.TotalValue())

	if combinedQuantity.IsZero() {
		return NewZeroMonetaryValue()
	}

	averageCostValue := combinedTotalCost.ToFloat64() / float64(combinedQuantity.ToInt())

	return NewMonetaryValue(averageCostValue)
}
