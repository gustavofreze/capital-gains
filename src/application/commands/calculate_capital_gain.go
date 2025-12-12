package commands

var _ Command = (*CalculateCapitalGain)(nil)

type CalculateCapitalGain struct{}

func NewCalculateCapitalGain() Command {
	return CalculateCapitalGain{}
}
