package models

type UnitCost struct {
	value MonetaryValue
}

func NewUnitCost(value float64) UnitCost {
	return UnitCost{
		value: NewMonetaryValue(value),
	}
}

func (unitCost UnitCost) Value() MonetaryValue {
	return unitCost.value
}

func (unitCost UnitCost) Subtract(other UnitCost) UnitCost {
	value := unitCost.value.Subtract(other.value)

	return NewUnitCost(value.ToFloat64())
}

func (unitCost UnitCost) MultiplyBy(quantity Quantity) MonetaryValue {
	return unitCost.value.MultiplyBy(quantity.ToFloat64())
}
