package commands

var _ Command = (*CalculateCapitalGain)(nil)

type CalculateCapitalGain struct{}

func NewCalculateCapitalGain() CalculateCapitalGain {
	return CalculateCapitalGain{}
}
