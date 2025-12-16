package driver

import "fmt"

type Tax struct {
	Value float64
}

func NewTax(value float64) Tax {
	return Tax{Value: value}
}

func (tax Tax) MarshalJSON() ([]byte, error) {
	jsonObject := fmt.Sprintf("{\"tax\":%.2f}", tax.Value)
	return []byte(jsonObject), nil
}
