package models

type Tax struct {
	value MonetaryValue
}

func NewTax(value MonetaryValue) Tax {
	return Tax{value: value}
}

func (tax Tax) Value() MonetaryValue {
	return tax.value
}

func (tax Tax) IsExempted() bool {
	return tax.value.IsZero()
}
