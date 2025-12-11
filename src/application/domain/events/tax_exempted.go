package events

type TaxExempted struct {
	amount float64
}

func NewTaxExempted() TaxExempted {
	return TaxExempted{
		amount: 0.00,
	}
}

func (tax TaxExempted) Amount() float64 {
	return tax.amount
}
