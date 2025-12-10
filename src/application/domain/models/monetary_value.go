package models

type MonetaryValue float64

func NewMonetaryValue(value float64) MonetaryValue {
	return MonetaryValue(value)
}

func (monetaryValue MonetaryValue) Add(other MonetaryValue) MonetaryValue {
	return monetaryValue + other
}

func (monetaryValue MonetaryValue) Subtract(other MonetaryValue) MonetaryValue {
	return monetaryValue - other
}

func (monetaryValue MonetaryValue) MultiplyBy(factor float64) MonetaryValue {
	return MonetaryValue(float64(monetaryValue) * factor)
}

func (monetaryValue MonetaryValue) IsZero() bool {
	return monetaryValue == 0
}

func (monetaryValue MonetaryValue) IsPositive() bool {
	return monetaryValue > 0
}

func (monetaryValue MonetaryValue) IsNegative() bool {
	return monetaryValue < 0
}

func (monetaryValue MonetaryValue) IsLessThan(other MonetaryValue) bool {
	return monetaryValue < other
}

func (monetaryValue MonetaryValue) IsGreaterThan(other MonetaryValue) bool {
	return monetaryValue > other
}

func (monetaryValue MonetaryValue) IsGreaterThanOrEqual(other MonetaryValue) bool {
	return monetaryValue >= other
}

func (monetaryValue MonetaryValue) AbsoluteValue() MonetaryValue {
	if monetaryValue.IsNegative() {
		return NewMonetaryValue(-float64(monetaryValue))
	}

	return monetaryValue
}

func (monetaryValue MonetaryValue) ToFloat64() float64 {
	return float64(monetaryValue)
}
