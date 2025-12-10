package models

type Quantity int64

func NewQuantity(value int64) Quantity {
	return Quantity(value)
}

func (quantity Quantity) Add(other Quantity) Quantity {
	return quantity + other
}

func (quantity Quantity) Subtract(other Quantity) Quantity {
	return quantity - other
}

func (quantity Quantity) IsZero() bool {
	return quantity == 0
}

func (quantity Quantity) ToFloat64() float64 {
	return float64(quantity)
}
