package models

type Quantity int

func NewQuantity(value int) Quantity {
	return Quantity(value)
}

func (quantity Quantity) Add(other Quantity) Quantity {
	return quantity + other
}

func (quantity Quantity) Subtract(other Quantity) Quantity {
	return quantity - other
}

func (quantity Quantity) ToInt() int {
	return int(quantity)
}

func (quantity Quantity) IsZero() bool {
	return quantity == 0
}

func (quantity Quantity) ToFloat() float64 {
	return float64(quantity)
}
