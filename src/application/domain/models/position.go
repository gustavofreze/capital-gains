package models

const (
	taxFreeThreshold MonetaryValue = 20000.00
	taxRate          float64       = 0.20
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
	combinedQuantity := position.quantity.Add(quantity)

	if combinedQuantity.IsZero() {
		position.quantity = combinedQuantity
		position.averageUnitCost = NewZeroMonetaryValue()
		return
	}

	currentTotalCost := position.averageUnitCost.MultiplyBy(position.quantity.ToFloat())
	buyTotalCost := unitCost.MultiplyBy(quantity.ToFloat())
	combinedTotalCost := currentTotalCost.Add(buyTotalCost)
	averageUnitCost := combinedTotalCost.ToFloat64() / float64(combinedQuantity.ToInt())

	position.averageUnitCost = NewMonetaryValue(averageUnitCost)
	position.quantity = combinedQuantity
}

func (position *Position) Sell(quantity Quantity, unitCost MonetaryValue) Tax {
	proceeds := unitCost.MultiplyBy(quantity.ToFloat())
	grossCapitalGain := unitCost.Subtract(position.averageUnitCost).MultiplyBy(quantity.ToFloat())

	position.quantity = position.quantity.Subtract(quantity)

	if position.quantity.IsZero() {
		position.averageUnitCost = NewZeroMonetaryValue()
	}

	if grossCapitalGain.IsNegative() {
		position.accumulatedLoss = position.accumulatedLoss.Add(grossCapitalGain.AbsoluteValue())
		return NewTax(NewZeroMonetaryValue())
	}

	if proceeds.IsGreaterThan(taxFreeThreshold) {
		if grossCapitalGain.IsZero() {
			return NewTax(NewZeroMonetaryValue())
		}

		netProfit := grossCapitalGain

		if position.accumulatedLoss.IsGreaterThanOrEqual(netProfit) {
			position.accumulatedLoss = position.accumulatedLoss.Subtract(netProfit)
			return NewTax(NewZeroMonetaryValue())
		}

		if position.accumulatedLoss.IsPositive() {
			netProfit = netProfit.Subtract(position.accumulatedLoss)
			position.accumulatedLoss = NewZeroMonetaryValue()
		}

		return NewTax(netProfit.MultiplyBy(taxRate))
	}

	return NewTax(NewZeroMonetaryValue())
}
