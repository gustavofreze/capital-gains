package models

import "math"

const zero float64 = 0.00

type MonetaryValue float64

func NewMonetaryValue(value float64) MonetaryValue {
	return MonetaryValue(value)
}

func NewZeroMonetaryValue() MonetaryValue {
	return MonetaryValue(zero)
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
	return monetaryValue.ToFloat64() == zero
}

func (monetaryValue MonetaryValue) IsPositive() bool {
	return monetaryValue.ToFloat64() > zero
}

func (monetaryValue MonetaryValue) IsNegative() bool {
	return monetaryValue.ToFloat64() < zero
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
	return NewMonetaryValue(math.Abs(monetaryValue.ToFloat64()))
}

func (monetaryValue MonetaryValue) ToFloat64() float64 {
	return float64(monetaryValue)
}
