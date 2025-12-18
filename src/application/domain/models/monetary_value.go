package models

import "math"

type MonetaryValue float64

const (
	monetaryScale     float64       = 100.00
	zeroMonetaryValue MonetaryValue = 0.00
)

func NewMonetaryValue(value float64) MonetaryValue {
	return MonetaryValue(roundToTwoDecimalPlaces(value))
}

func NewZeroMonetaryValue() MonetaryValue {
	return zeroMonetaryValue
}

func (monetaryValue MonetaryValue) Add(other MonetaryValue) MonetaryValue {
	return NewMonetaryValue(monetaryValue.ToFloat64() + other.ToFloat64())
}

func (monetaryValue MonetaryValue) Subtract(other MonetaryValue) MonetaryValue {
	return NewMonetaryValue(monetaryValue.ToFloat64() - other.ToFloat64())
}

func (monetaryValue MonetaryValue) MultiplyBy(factor float64) MonetaryValue {
	return NewMonetaryValue(monetaryValue.ToFloat64() * factor)
}

func (monetaryValue MonetaryValue) IsZero() bool {
	return monetaryValue == zeroMonetaryValue
}

func (monetaryValue MonetaryValue) IsPositive() bool {
	return monetaryValue > 0
}

func (monetaryValue MonetaryValue) IsNegative() bool {
	return monetaryValue < 0
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
	return roundToTwoDecimalPlaces(float64(monetaryValue))
}

func roundToTwoDecimalPlaces(value float64) float64 {
	return math.Round(value*monetaryScale) / monetaryScale
}
