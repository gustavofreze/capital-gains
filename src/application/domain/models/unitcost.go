package models

type UnitCost struct {
	value float64
}

func NewUnitCost(value float64) UnitCost {
	return UnitCost{value: value}
}

func (unitCost UnitCost) Subtract(other UnitCost) UnitCost {
	return NewUnitCost(unitCost.value - other.value)
}

func (unitCost UnitCost) MultiplyBy(quantity Quantity) MonetaryValue {
	return NewMonetaryValue(unitCost.value * quantity.ToFloat64())
}
