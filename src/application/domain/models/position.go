package models

const (
	TaxRate                        = 0.20
	TaxFreeThreshold MonetaryValue = 20000.00
)

type Position struct {
	quantity        Quantity
	averageUnitCost MonetaryValue
	accumulatedLoss MonetaryValue
}

func NewPosition() Position {
	return Position{
		quantity:        NewQuantity(0),
		averageUnitCost: NewZeroMonetaryValue(),
		accumulatedLoss: NewZeroMonetaryValue(),
	}
}

func (position *Position) Buy(quantity Quantity, unitCost MonetaryValue) {
	position.averageUnitCost = position.calculateWeightedAverageUnitCost(quantity, unitCost)
	position.quantity = position.quantity.Add(quantity)
}

func (position *Position) Sell(quantity Quantity, unitCost MonetaryValue) MonetaryValue {
	totalProceeds := unitCost.MultiplyBy(quantity.ToFloat())
	grossCapitalGain := unitCost.
		Subtract(position.averageUnitCost).
		MultiplyBy(quantity.ToFloat())

	position.updateAccumulatedLossFromGrossGain(grossCapitalGain)

	netCapitalGain := position.calculateNetCapitalGain(grossCapitalGain)

	taxAmount := NewZeroMonetaryValue()
	hasTaxableProceeds := totalProceeds.IsGreaterThan(TaxFreeThreshold)
	hasTaxableGain := netCapitalGain.IsPositive()

	if hasTaxableProceeds && hasTaxableGain {
		taxAmount = netCapitalGain.MultiplyBy(TaxRate)
	}

	position.quantity = position.quantity.Subtract(quantity)

	if position.quantity.IsZero() {
		position.averageUnitCost = NewZeroMonetaryValue()
	}

	return taxAmount
}

func (position *Position) calculateWeightedAverageUnitCost(
	buyQuantity Quantity,
	buyUnitCost MonetaryValue,
) MonetaryValue {
	currentTotalCost := position.averageUnitCost.MultiplyBy(position.quantity.ToFloat())
	buyTotalCost := buyUnitCost.MultiplyBy(buyQuantity.ToFloat())

	combinedQuantity := position.quantity.Add(buyQuantity)
	if combinedQuantity.IsZero() {
		return NewZeroMonetaryValue()
	}

	combinedTotalCost := currentTotalCost.Add(buyTotalCost)
	averageCostValue := combinedTotalCost.ToFloat64() / float64(combinedQuantity.ToInt())

	return NewMonetaryValue(averageCostValue)
}

func (position *Position) updateAccumulatedLossFromGrossGain(grossCapitalGain MonetaryValue) {
	if grossCapitalGain.IsNegative() {
		position.accumulatedLoss = position.accumulatedLoss.Add(grossCapitalGain.AbsoluteValue())
	}
}

func (position *Position) calculateNetCapitalGain(grossCapitalGain MonetaryValue) MonetaryValue {
	if !grossCapitalGain.IsPositive() {
		return NewZeroMonetaryValue()
	}

	netCapitalGain := grossCapitalGain

	if position.accumulatedLoss.IsGreaterThanOrEqual(netCapitalGain) {
		position.accumulatedLoss = position.accumulatedLoss.Subtract(netCapitalGain)
		return NewZeroMonetaryValue()
	}

	if position.accumulatedLoss.IsPositive() {
		netCapitalGain = netCapitalGain.Subtract(position.accumulatedLoss)
		position.accumulatedLoss = NewZeroMonetaryValue()
	}

	return netCapitalGain
}
