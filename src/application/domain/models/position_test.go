package models_test

import (
	"testing"

	"capital-gains/src/application/domain/models"

	"github.com/stretchr/testify/assert"
)

func TestPositionApplyBuyGivenSingleBuyWhenApplyBuyThenPositionIsUpdated(t *testing.T) {
	t.Parallel()

	// Given I have an empty position
	position := models.NewEmptyPosition()

	// And a unit cost of 10.0
	unitCost := models.NewUnitCost(10.0)

	// And a quantity of 100
	quantity := models.NewQuantity(100)

	// When I create a buy operation with the unit cost and quantity
	buyOperation := models.NewBuy(unitCost, quantity)

	// And I apply the buy operation to the position
	position.ApplyBuy(buyOperation)

	// Then the position quantity should be equal to the buy quantity
	assert.Equal(t, quantity, position.Quantity())

	// And the position average unit cost should be equal to the buy unit cost
	assert.Equal(t, unitCost, position.AverageUnitCost())

	// And the position accumulated loss should be zero
	assert.True(t, position.AccumulatedLoss().IsZero())
}

func TestPositionApplyBuyGivenZeroQuantityBuyWhenApplyBuyThenPositionRemainsUnchanged(t *testing.T) {
	t.Parallel()

	// Given I have an empty position
	position := models.NewEmptyPosition()

	// And a unit cost of 10.0 for the zero quantity buy
	unitCost := models.NewUnitCost(10.0)

	// And a zero buy quantity
	quantity := models.NewQuantity(0)

	// When I create a buy operation with the unit cost and zero quantity
	buyOperation := models.NewBuy(unitCost, quantity)

	// And I apply the buy operation to the position
	position.ApplyBuy(buyOperation)

	// Then the position quantity should remain zero
	assert.True(t, position.Quantity().IsZero())

	// And the position average unit cost should remain zero
	assert.Equal(t, models.NewUnitCost(0), position.AverageUnitCost())

	// And the position accumulated loss should remain zero
	assert.True(t, position.AccumulatedLoss().IsZero())
}

func TestPositionApplySellGivenProfitsWithProceedsBelowThresholdWhenApplySellThenNoTaxIsCharged(t *testing.T) {
	t.Parallel()

	// Given I have an empty position
	position := models.NewEmptyPosition()

	// And a buy unit cost of 10.0
	buyUnitCost := models.NewUnitCost(10.0)

	// And a buy quantity of 100
	buyQuantity := models.NewQuantity(100)

	// And I create a buy operation with the buy unit cost and buy quantity
	buyOperation := models.NewBuy(buyUnitCost, buyQuantity)

	// And I apply the buy operation to the position
	position.ApplyBuy(buyOperation)

	// And a first sell unit cost of 15.0
	firstSellUnitCost := models.NewUnitCost(15.0)

	// And a first sell quantity of 50
	firstSellQuantity := models.NewQuantity(50)

	// And I create the first sell operation with the first sell unit cost and first sell quantity
	firstSellOperation := models.NewSell(firstSellUnitCost, firstSellQuantity)

	// When I apply the first sell operation to the position
	firstSellTax := position.ApplySell(firstSellOperation)

	// And a second sell unit cost of 15.0
	secondSellUnitCost := models.NewUnitCost(15.0)

	// And a second sell quantity of 50
	secondSellQuantity := models.NewQuantity(50)

	// And I create the second sell operation with the second sell unit cost and second sell quantity
	secondSellOperation := models.NewSell(secondSellUnitCost, secondSellQuantity)

	// And I apply the second sell operation to the position
	secondSellTax := position.ApplySell(secondSellOperation)

	// Then the first sell tax should be zero
	assert.True(t, firstSellTax.IsZero())

	// And the second sell tax should be zero
	assert.True(t, secondSellTax.IsZero())

	// And the position quantity should be zero
	assert.True(t, position.Quantity().IsZero())

	// And the position average unit cost should be zero
	assert.Equal(t, models.NewUnitCost(0), position.AverageUnitCost())

	// And the position accumulated loss should be zero
	assert.True(t, position.AccumulatedLoss().IsZero())
}

func TestPositionApplySellGivenHighProfitThenSubsequentLossWhenApplySellThenOnlyFirstProfitIsTaxed(t *testing.T) {
	t.Parallel()

	// Given I have an empty position
	position := models.NewEmptyPosition()

	// And a buy unit cost of 10.0
	buyUnitCost := models.NewUnitCost(10.0)

	// And a buy quantity of 10000
	buyQuantity := models.NewQuantity(10000)

	// And I create a buy operation with the buy unit cost and buy quantity
	buyOperation := models.NewBuy(buyUnitCost, buyQuantity)

	// And I apply the buy operation to the position
	position.ApplyBuy(buyOperation)

	// And a first sell unit cost of 20.0
	firstSellUnitCost := models.NewUnitCost(20.0)

	// And a first sell quantity of 5000
	firstSellQuantity := models.NewQuantity(5000)

	// And I create the first sell operation with the first sell unit cost and first sell quantity
	firstSellOperation := models.NewSell(firstSellUnitCost, firstSellQuantity)

	// When I apply the first sell operation to the position
	firstSellTax := position.ApplySell(firstSellOperation)

	// And a second sell unit cost of 5.0
	secondSellUnitCost := models.NewUnitCost(5.0)

	// And a second sell quantity of 5000
	secondSellQuantity := models.NewQuantity(5000)

	// And I create the second sell operation with the second sell unit cost and second sell quantity
	secondSellOperation := models.NewSell(secondSellUnitCost, secondSellQuantity)

	// And I apply the second sell operation to the position
	secondSellTax := position.ApplySell(secondSellOperation)

	// Then the first sell tax should be equal to 10000.0
	assert.Equal(t, models.NewMonetaryValue(10000.0), firstSellTax)

	// And the second sell tax should be zero
	assert.True(t, secondSellTax.IsZero())

	// And after the second sell the accumulated loss should be 25000.0
	assert.Equal(t, models.NewMonetaryValue(25000.0), position.AccumulatedLoss())
}

func TestPositionApplySellGivenLossThenProfitWhenApplySellThenLossIsDeductedFromProfitBeforeTax(t *testing.T) {
	t.Parallel()

	// Given I have an empty position
	position := models.NewEmptyPosition()

	// And a buy unit cost of 10.0
	buyUnitCost := models.NewUnitCost(10.0)

	// And a buy quantity of 10000
	buyQuantity := models.NewQuantity(10000)

	// And I create a buy operation with the buy unit cost and buy quantity
	buyOperation := models.NewBuy(buyUnitCost, buyQuantity)

	// And I apply the buy operation to the position
	position.ApplyBuy(buyOperation)

	// And a first sell unit cost of 5.0
	firstSellUnitCost := models.NewUnitCost(5.0)

	// And a first sell quantity of 5000
	firstSellQuantity := models.NewQuantity(5000)

	// And I create the first sell operation with the first sell unit cost and first sell quantity
	firstSellOperation := models.NewSell(firstSellUnitCost, firstSellQuantity)

	// When I apply the first sell operation to the position
	firstSellTax := position.ApplySell(firstSellOperation)

	// Then the first sell tax should be zero
	assert.True(t, firstSellTax.IsZero())

	// And after the first sell the accumulated loss should be 25000.0
	assert.Equal(t, models.NewMonetaryValue(25000.0), position.AccumulatedLoss())

	// And a second sell unit cost of 20.0
	secondSellUnitCost := models.NewUnitCost(20.0)

	// And a second sell quantity of 3000
	secondSellQuantity := models.NewQuantity(3000)

	// And I create the second sell operation with the second sell unit cost and second sell quantity
	secondSellOperation := models.NewSell(secondSellUnitCost, secondSellQuantity)

	// And I apply the second sell operation to the position
	secondSellTax := position.ApplySell(secondSellOperation)

	// Then the second sell tax should be equal to 1000.0
	assert.Equal(t, models.NewMonetaryValue(1000.0), secondSellTax)

	// And after the second sell the accumulated loss should be zero
	assert.True(t, position.AccumulatedLoss().IsZero())
}

func TestPositionApplySellGivenTwoBuysAndSellAtWeightedAverageWhenApplySellThenNoTaxIsCharged(t *testing.T) {
	t.Parallel()

	// Given I have an empty position
	position := models.NewEmptyPosition()

	// And a first buy unit cost of 10.0
	firstBuyUnitCost := models.NewUnitCost(10.0)

	// And a first buy quantity of 10000
	firstBuyQuantity := models.NewQuantity(10000)

	// And I create the first buy operation with the first buy unit cost and first buy quantity
	firstBuyOperation := models.NewBuy(firstBuyUnitCost, firstBuyQuantity)

	// And I apply the first buy operation to the position
	position.ApplyBuy(firstBuyOperation)

	// And a second buy unit cost of 25.0
	secondBuyUnitCost := models.NewUnitCost(25.0)

	// And a second buy quantity of 5000
	secondBuyQuantity := models.NewQuantity(5000)

	// And I create the second buy operation with the second buy unit cost and second buy quantity
	secondBuyOperation := models.NewBuy(secondBuyUnitCost, secondBuyQuantity)

	// And I apply the second buy operation to the position
	position.ApplyBuy(secondBuyOperation)

	// And I expect the weighted average unit cost to be 15.0 before any sell
	expectedAverageUnitCostBeforeSell := models.NewUnitCost(15.0)
	assert.Equal(t, expectedAverageUnitCostBeforeSell, position.AverageUnitCost())

	// And a sell unit cost of 15.0
	sellUnitCost := models.NewUnitCost(15.0)

	// And a sell quantity of 10000
	sellQuantity := models.NewQuantity(10000)

	// And I create the sell operation with the sell unit cost and sell quantity
	sellOperation := models.NewSell(sellUnitCost, sellQuantity)

	// When I apply the sell operation to the position
	sellTax := position.ApplySell(sellOperation)

	// Then the sell tax should be zero
	assert.True(t, sellTax.IsZero())

	// And the position quantity should be 5000
	assert.Equal(t, models.NewQuantity(5000), position.Quantity())

	// And the position average unit cost should remain 15.0
	assert.Equal(t, models.NewUnitCost(15.0), position.AverageUnitCost())

	// And the position accumulated loss should be zero
	assert.True(t, position.AccumulatedLoss().IsZero())
}

func TestPositionApplySellGivenSellAtAverageThenSellAboveAverageWhenApplySellThenOnlySecondSellIsTaxed(t *testing.T) {
	t.Parallel()

	// Given I have an empty position
	position := models.NewEmptyPosition()

	// And a first buy unit cost of 10.0
	firstBuyUnitCost := models.NewUnitCost(10.0)

	// And a first buy quantity of 10000
	firstBuyQuantity := models.NewQuantity(10000)

	// And I create the first buy operation with the first buy unit cost and first buy quantity
	firstBuyOperation := models.NewBuy(firstBuyUnitCost, firstBuyQuantity)

	// And I apply the first buy operation to the position
	position.ApplyBuy(firstBuyOperation)

	// And a second buy unit cost of 25.0
	secondBuyUnitCost := models.NewUnitCost(25.0)

	// And a second buy quantity of 5000
	secondBuyQuantity := models.NewQuantity(5000)

	// And I create the second buy operation with the second buy unit cost and second buy quantity
	secondBuyOperation := models.NewBuy(secondBuyUnitCost, secondBuyQuantity)

	// And I apply the second buy operation to the position
	position.ApplyBuy(secondBuyOperation)

	// And a first sell unit cost of 15.0
	firstSellUnitCost := models.NewUnitCost(15.0)

	// And a first sell quantity of 10000
	firstSellQuantity := models.NewQuantity(10000)

	// And I create the first sell operation with the first sell unit cost and first sell quantity
	firstSellOperation := models.NewSell(firstSellUnitCost, firstSellQuantity)

	// When I apply the first sell operation to the position
	firstSellTax := position.ApplySell(firstSellOperation)

	// And a second sell unit cost of 25.0
	secondSellUnitCost := models.NewUnitCost(25.0)

	// And a second sell quantity of 5000
	secondSellQuantity := models.NewQuantity(5000)

	// And I create the second sell operation with the second sell unit cost and second sell quantity
	secondSellOperation := models.NewSell(secondSellUnitCost, secondSellQuantity)

	// And I apply the second sell operation to the position
	secondSellTax := position.ApplySell(secondSellOperation)

	// Then the first sell tax should be zero
	assert.True(t, firstSellTax.IsZero())

	// And the second sell tax should be equal to 10000.0
	assert.Equal(t, models.NewMonetaryValue(10000.0), secondSellTax)

	// And the position quantity should be zero
	assert.True(t, position.Quantity().IsZero())

	// And the position average unit cost should be zero
	assert.Equal(t, models.NewUnitCost(0), position.AverageUnitCost())

	// And the position accumulated loss should be zero
	assert.True(t, position.AccumulatedLoss().IsZero())
}

func TestPositionApplySellGivenLossBelowThresholdThenMultipleProfitsWhenApplySellThenLossAndThresholdRulesAreApplied(t *testing.T) {
	t.Parallel()

	// Given I have an empty position
	position := models.NewEmptyPosition()

	// And a buy unit cost of 10.0
	buyUnitCost := models.NewUnitCost(10.0)

	// And a buy quantity of 10000
	buyQuantity := models.NewQuantity(10000)

	// And I create a buy operation with the buy unit cost and buy quantity
	buyOperation := models.NewBuy(buyUnitCost, buyQuantity)

	// And I apply the buy operation to the position
	position.ApplyBuy(buyOperation)

	// And a first sell unit cost of 2.0
	firstSellUnitCost := models.NewUnitCost(2.0)

	// And a first sell quantity of 5000
	firstSellQuantity := models.NewQuantity(5000)

	// And I create the first sell operation with the first sell unit cost and first sell quantity
	firstSellOperation := models.NewSell(firstSellUnitCost, firstSellQuantity)

	// When I apply the first sell operation to the position
	firstSellTax := position.ApplySell(firstSellOperation)

	// Then the first sell tax should be zero
	assert.True(t, firstSellTax.IsZero())

	// And after the first sell the accumulated loss should be 40000.0
	expectedAccumulatedLossAfterFirstSell := models.NewMonetaryValue(40000.0)
	assert.Equal(t, expectedAccumulatedLossAfterFirstSell, position.AccumulatedLoss())

	// And a second sell unit cost of 20.0
	secondSellUnitCost := models.NewUnitCost(20.0)

	// And a second sell quantity of 2000
	secondSellQuantity := models.NewQuantity(2000)

	// And I create the second sell operation with the second sell unit cost and second sell quantity
	secondSellOperation := models.NewSell(secondSellUnitCost, secondSellQuantity)

	// When I apply the second sell operation to the position
	secondSellTax := position.ApplySell(secondSellOperation)

	// Then the second sell tax should be zero
	assert.True(t, secondSellTax.IsZero())

	// And after the second sell the accumulated loss should be 20000.0
	expectedAccumulatedLossAfterSecondSell := models.NewMonetaryValue(20000.0)
	assert.Equal(t, expectedAccumulatedLossAfterSecondSell, position.AccumulatedLoss())

	// And a third sell unit cost of 20.0
	thirdSellUnitCost := models.NewUnitCost(20.0)

	// And a third sell quantity of 2000
	thirdSellQuantity := models.NewQuantity(2000)

	// And I create the third sell operation with the third sell unit cost and third sell quantity
	thirdSellOperation := models.NewSell(thirdSellUnitCost, thirdSellQuantity)

	// When I apply the third sell operation to the position
	thirdSellTax := position.ApplySell(thirdSellOperation)

	// Then the third sell tax should be zero
	assert.True(t, thirdSellTax.IsZero())

	// And after the third sell the accumulated loss should be zero
	assert.True(t, position.AccumulatedLoss().IsZero())

	// And a fourth sell unit cost of 25.0
	fourthSellUnitCost := models.NewUnitCost(25.0)

	// And a fourth sell quantity of 1000
	fourthSellQuantity := models.NewQuantity(1000)

	// And I create the fourth sell operation with the fourth sell unit cost and fourth sell quantity
	fourthSellOperation := models.NewSell(fourthSellUnitCost, fourthSellQuantity)

	// When I apply the fourth sell operation to the position
	fourthSellTax := position.ApplySell(fourthSellOperation)

	// Then the fourth sell tax should be equal to 3000.0
	assert.Equal(t, models.NewMonetaryValue(3000.0), fourthSellTax)

	// And the position remaining quantity should be 1000
	assert.Equal(t, models.NewQuantity(0), position.Quantity())
}

func TestPositionApplySellGivenMultipleCyclesWithLossCompensationAndThresholdWhenApplySellThenTaxesMatchSpecification(t *testing.T) {
	t.Parallel()

	// Given I have an empty position
	position := models.NewEmptyPosition()

	// And a first cycle buy unit cost of 10.0
	firstCycleBuyUnitCost := models.NewUnitCost(10.0)

	// And a first cycle buy quantity of 10000
	firstCycleBuyQuantity := models.NewQuantity(10000)

	// And I create a first cycle buy operation with the first cycle buy unit cost and first cycle buy quantity
	firstCycleBuyOperation := models.NewBuy(firstCycleBuyUnitCost, firstCycleBuyQuantity)

	// And I apply the first cycle buy operation to the position
	position.ApplyBuy(firstCycleBuyOperation)

	// And a first cycle first sell unit cost of 2.0
	firstCycleFirstSellUnitCost := models.NewUnitCost(2.0)

	// And a first cycle first sell quantity of 5000
	firstCycleFirstSellQuantity := models.NewQuantity(5000)

	// And I create a first cycle first sell operation with the first cycle first sell unit cost and first cycle sell quantity
	firstCycleFirstSellOperation := models.NewSell(firstCycleFirstSellUnitCost, firstCycleFirstSellQuantity)

	// When I apply the first cycle first sell operation to the position
	firstCycleFirstSellTax := position.ApplySell(firstCycleFirstSellOperation)

	// And a first cycle second sell unit cost of 20.0
	firstCycleSecondSellUnitCost := models.NewUnitCost(20.0)

	// And a first cycle second sell quantity of 2000
	firstCycleSecondSellQuantity := models.NewQuantity(2000)

	// And I create a first cycle second sell operation with the first cycle second sell unit cost and first cycle second sell quantity
	firstCycleSecondSellOperation := models.NewSell(firstCycleSecondSellUnitCost, firstCycleSecondSellQuantity)

	// And I apply the first cycle second sell operation to the position
	firstCycleSecondSellTax := position.ApplySell(firstCycleSecondSellOperation)

	// And a first cycle third sell unit cost of 20.0
	firstCycleThirdSellUnitCost := models.NewUnitCost(20.0)

	// And a first cycle third sell quantity of 2000
	firstCycleThirdSellQuantity := models.NewQuantity(2000)

	// And I create a first cycle third sell operation with the first cycle third sell unit cost and first cycle third sell quantity
	firstCycleThirdSellOperation := models.NewSell(firstCycleThirdSellUnitCost, firstCycleThirdSellQuantity)

	// And I apply the first cycle third sell operation to the position
	firstCycleThirdSellTax := position.ApplySell(firstCycleThirdSellOperation)

	// And a first cycle fourth sell unit cost of 25.0
	firstCycleFourthSellUnitCost := models.NewUnitCost(25.0)

	// And a first cycle fourth sell quantity of 1000
	firstCycleFourthSellQuantity := models.NewQuantity(1000)

	// And I create a first cycle fourth sell operation with the first cycle fourth sell unit cost and first cycle fourth sell quantity
	firstCycleFourthSellOperation := models.NewSell(firstCycleFourthSellUnitCost, firstCycleFourthSellQuantity)

	// And I apply the first cycle fourth sell operation to the position
	firstCycleFourthSellTax := position.ApplySell(firstCycleFourthSellOperation)

	// And a second cycle buy unit cost of 20.0
	secondCycleBuyUnitCost := models.NewUnitCost(20.0)

	// And a second cycle buy quantity of 10000
	secondCycleBuyQuantity := models.NewQuantity(10000)

	// And I create a second cycle buy operation with the second cycle buy unit cost and second cycle buy quantity
	secondCycleBuyOperation := models.NewBuy(secondCycleBuyUnitCost, secondCycleBuyQuantity)

	// And I apply the second cycle buy operation to the position
	position.ApplyBuy(secondCycleBuyOperation)

	// And a second cycle first sell unit cost of 15.0
	secondCycleFirstSellUnitCost := models.NewUnitCost(15.0)

	// And a second cycle first sell quantity of 5000
	secondCycleFirstSellQuantity := models.NewQuantity(5000)

	// And I create a second cycle first sell operation with the second cycle first sell unit cost and second cycle first sell quantity
	secondCycleFirstSellOperation := models.NewSell(secondCycleFirstSellUnitCost, secondCycleFirstSellQuantity)

	// And I apply the second cycle first sell operation to the position
	secondCycleFirstSellTax := position.ApplySell(secondCycleFirstSellOperation)

	// And a second cycle second sell unit cost of 30.0
	secondCycleSecondSellUnitCost := models.NewUnitCost(30.0)

	// And a second cycle second sell quantity of 4350
	secondCycleSecondSellQuantity := models.NewQuantity(4350)

	// And I create a second cycle second sell operation with the second cycle second sell unit cost and second cycle second sell quantity
	secondCycleSecondSellOperation := models.NewSell(secondCycleSecondSellUnitCost, secondCycleSecondSellQuantity)

	// And I apply the second cycle second sell operation to the position
	secondCycleSecondSellTax := position.ApplySell(secondCycleSecondSellOperation)

	// And a second cycle third sell unit cost of 30.0
	secondCycleThirdSellUnitCost := models.NewUnitCost(30.0)

	// And a second cycle third sell quantity of 650
	secondCycleThirdSellQuantity := models.NewQuantity(650)

	// And I create a second cycle third sell operation with the second cycle third sell unit cost and second cycle third sell quantity
	secondCycleThirdSellOperation := models.NewSell(secondCycleThirdSellUnitCost, secondCycleThirdSellQuantity)

	// And I apply the second cycle third sell operation to the position
	secondCycleThirdSellTax := position.ApplySell(secondCycleThirdSellOperation)

	// Then the first cycle first sell tax should be zero
	assert.True(t, firstCycleFirstSellTax.IsZero())

	// And the first cycle second sell tax should be zero
	assert.True(t, firstCycleSecondSellTax.IsZero())

	// And the first cycle third sell tax should be zero
	assert.True(t, firstCycleThirdSellTax.IsZero())

	// And the first cycle fourth sell tax should be equal to 3000.0
	assert.Equal(t, models.NewMonetaryValue(3000.0), firstCycleFourthSellTax)

	// And the second cycle first sell tax should be zero
	assert.True(t, secondCycleFirstSellTax.IsZero())

	// And the second cycle second sell tax should be equal to 3700.0
	assert.Equal(t, models.NewMonetaryValue(3700.0), secondCycleSecondSellTax)

	// And the second cycle third sell tax should be zero
	assert.True(t, secondCycleThirdSellTax.IsZero())
}

func TestPositionApplySellGivenTwoIndependentHighProfitCyclesWhenApplySellThenTaxIsCalculatedForEachCycle(t *testing.T) {
	t.Parallel()

	// Given I have an empty position
	position := models.NewEmptyPosition()

	// And a first cycle buy unit cost of 10.0
	firstCycleBuyUnitCost := models.NewUnitCost(10.0)

	// And a first cycle buy quantity of 10000
	firstCycleBuyQuantity := models.NewQuantity(10000)

	// And I create a first cycle buy operation with the first cycle buy unit cost and first cycle buy quantity
	firstCycleBuyOperation := models.NewBuy(firstCycleBuyUnitCost, firstCycleBuyQuantity)

	// And I apply the first cycle buy operation to the position
	position.ApplyBuy(firstCycleBuyOperation)

	// And a first cycle sell unit cost of 50.0
	firstCycleSellUnitCost := models.NewUnitCost(50.0)

	// And a first cycle sell quantity of 10000
	firstCycleSellQuantity := models.NewQuantity(10000)

	// And I create a first cycle sell operation with the first cycle sell unit cost and first cycle sell quantity
	firstCycleSellOperation := models.NewSell(firstCycleSellUnitCost, firstCycleSellQuantity)

	// When I apply the first cycle sell operation to the position
	firstCycleSellTax := position.ApplySell(firstCycleSellOperation)

	// And a second cycle buy unit cost of 20.0
	secondCycleBuyUnitCost := models.NewUnitCost(20.0)

	// And a second cycle buy quantity of 10000
	secondCycleBuyQuantity := models.NewQuantity(10000)

	// And I create a second cycle buy operation with the second cycle buy unit cost and second cycle buy quantity
	secondCycleBuyOperation := models.NewBuy(secondCycleBuyUnitCost, secondCycleBuyQuantity)

	// And I apply the second cycle buy operation to the position
	position.ApplyBuy(secondCycleBuyOperation)

	// And a second cycle sell unit cost of 50.0
	secondCycleSellUnitCost := models.NewUnitCost(50.0)

	// And a second cycle sell quantity of 10000
	secondCycleSellQuantity := models.NewQuantity(10000)

	// And I create a second cycle sell operation with the second cycle sell unit cost and second cycle sell quantity
	secondCycleSellOperation := models.NewSell(secondCycleSellUnitCost, secondCycleSellQuantity)

	// And I apply the second cycle sell operation to the position
	secondCycleSellTax := position.ApplySell(secondCycleSellOperation)

	// Then the first cycle sell tax should be equal to 80000.0
	assert.Equal(t, models.NewMonetaryValue(80000.0), firstCycleSellTax)

	// And the second cycle sell tax should be equal to 60000.0
	assert.Equal(t, models.NewMonetaryValue(60000.0), secondCycleSellTax)
}

func TestPositionApplySellGivenSellAtAverageUnitCostWhenApplySellThenNoProfitNoLossAndNoTaxIsCharged(t *testing.T) {
	t.Parallel()

	// Given I have an empty position
	position := models.NewEmptyPosition()

	// And a buy unit cost of 10.0
	buyUnitCost := models.NewUnitCost(10.0)

	// And a buy quantity of 10000
	buyQuantity := models.NewQuantity(10000)

	// And I create a buy operation with the buy unit cost and buy quantity
	buyOperation := models.NewBuy(buyUnitCost, buyQuantity)

	// And I apply the buy operation to the position
	position.ApplyBuy(buyOperation)

	// And I expect the average unit cost to be equal to the buy unit cost
	assert.Equal(t, buyUnitCost, position.AverageUnitCost())

	// And a sell unit cost equal to the average unit cost
	sellUnitCost := position.AverageUnitCost()

	// And a sell quantity equal to the buy quantity
	sellQuantity := buyQuantity

	// And I create a sell operation with the sell unit cost and sell quantity
	sellOperation := models.NewSell(sellUnitCost, sellQuantity)

	// When I apply the sell operation to the position
	sellTax := position.ApplySell(sellOperation)

	// Then the sell tax should be zero
	assert.True(t, sellTax.IsZero())

	// And the position accumulated loss should be zero
	assert.True(t, position.AccumulatedLoss().IsZero())

	// And the position quantity should be zero
	assert.True(t, position.Quantity().IsZero())

	// And the position average unit cost should be reset to zero
	assert.Equal(t, models.NewUnitCost(0), position.AverageUnitCost())
}
