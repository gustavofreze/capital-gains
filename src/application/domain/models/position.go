package models

const TaxRate = 0.20
const TaxFreeThreshold MonetaryValue = 20000.00

type Position struct {
	quantity        Quantity
	averageUnitCost UnitCost
	accumulatedLoss MonetaryValue
}

func NewPosition() Position {
	return Position{
		quantity:        NewQuantity(0),
		averageUnitCost: NewUnitCost(0.00),
		accumulatedLoss: NewZeroMonetaryValue(),
	}
}

func (position *Position) ApplyBuy(buy Buy) {
	position.averageUnitCost = buy.CalculateWeightedAverageUnitCost(position.quantity, position.averageUnitCost)
	position.quantity = position.quantity.Add(buy.Quantity())
}

func (position *Position) ApplySell(sell Sell) MonetaryValue {
	totalProceeds := sell.TotalProceeds()
	grossCapitalGain := sell.CalculateGrossCapitalGain(position.averageUnitCost)

	if grossCapitalGain.IsNegative() {
		lossAsPositiveValue := grossCapitalGain.AbsoluteValue()
		position.accumulatedLoss = position.accumulatedLoss.Add(lossAsPositiveValue)
	}

	netCapitalGain := NewZeroMonetaryValue()

	if grossCapitalGain.IsPositive() {
		netCapitalGain = grossCapitalGain

		if position.accumulatedLoss.IsGreaterThanOrEqual(netCapitalGain) {
			position.accumulatedLoss = position.accumulatedLoss.Subtract(netCapitalGain)
			netCapitalGain = NewZeroMonetaryValue()
		}

		if netCapitalGain.IsPositive() && position.accumulatedLoss.IsLessThan(netCapitalGain) {
			netCapitalGain = netCapitalGain.Subtract(position.accumulatedLoss)
			position.accumulatedLoss = NewZeroMonetaryValue()
		}
	}

	taxAmount := NewZeroMonetaryValue()
	hasTaxableProceeds := totalProceeds.IsGreaterThan(TaxFreeThreshold)
	hasTaxableGain := netCapitalGain.IsPositive()

	if hasTaxableProceeds && hasTaxableGain {
		taxAmount = netCapitalGain.MultiplyBy(TaxRate)
	}

	position.quantity = position.quantity.Subtract(sell.Quantity())

	if position.quantity.IsZero() {
		position.averageUnitCost = NewUnitCost(0)
	}

	return taxAmount
}

func (position *Position) Quantity() Quantity {
	return position.quantity
}

func (position *Position) AverageUnitCost() UnitCost {
	return position.averageUnitCost
}

func (position *Position) AccumulatedLoss() MonetaryValue {
	return position.accumulatedLoss
}
