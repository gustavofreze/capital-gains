package events

type TaxPaid struct {
	amount float64
}

func NewTaxPaid(amount float64) TaxPaid {
	return TaxPaid{
		amount: amount,
	}
}

func (tax TaxPaid) Amount() float64 {
	return tax.amount
}
