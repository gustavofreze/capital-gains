package models

type Buy struct {
	unitCost UnitCost
	quantity Quantity
}

func NewBuy(unitCost UnitCost, quantity Quantity) Buy {
	return Buy{
		unitCost: unitCost,
		quantity: quantity,
	}
}

func (buy Buy) Quantity() Quantity {
	return buy.quantity
}

func (buy Buy) TotalCost() MonetaryValue {
	return buy.unitCost.MultiplyBy(buy.quantity)
}

func (buy Buy) CalculateWeightedAverageUnitCost(existingQuantity Quantity, existingUnitCost UnitCost) UnitCost {
	newTotalCost := buy.TotalCost()
	combinedQuantity := existingQuantity.Add(buy.quantity)
	existingTotalCost := existingUnitCost.MultiplyBy(existingQuantity)
	combinedTotalCost := existingTotalCost.Add(newTotalCost)

	if combinedQuantity.IsZero() {
		return NewUnitCost(0)
	}

	averageCostValue := combinedTotalCost.ToFloat64() / combinedQuantity.ToFloat64()
	return NewUnitCost(averageCostValue)
}
