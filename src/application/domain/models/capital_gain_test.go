package models_test

import (
	"testing"

	"capital-gains/src/application/domain/events"
	"capital-gains/src/application/domain/models"

	"github.com/stretchr/testify/assert"
)

func TestCapitalGainsApplyOperationsGivenProfitsWithProceedsBelowThresholdWhenApplyOperationsThenNoTaxIsCharged(t *testing.T) {
	t.Parallel()

	// Given I have an empty capital gains aggregate
	capitalGains := models.NewCapitalGains()

	// And a buy unit cost of 10.0
	buyUnitCost := models.NewUnitCost(10.0)

	// And a buy quantity of 100
	buyQuantity := models.NewQuantity(100)

	// And I create a buy operation with the buy unit cost and buy quantity
	buyOperation := models.NewBuy(buyUnitCost, buyQuantity)

	// And a first sell unit cost of 15.0
	firstSellUnitCost := models.NewUnitCost(15.0)

	// And a first sell quantity of 50
	firstSellQuantity := models.NewQuantity(50)

	// And I create the first sell operation with the first sell unit cost and first sell quantity
	firstSellOperation := models.NewSell(firstSellUnitCost, firstSellQuantity)

	// And a second sell unit cost of 15.0
	secondSellUnitCost := models.NewUnitCost(15.0)

	// And a second sell quantity of 50
	secondSellQuantity := models.NewQuantity(50)

	// And I create the second sell operation with the second sell unit cost and second sell quantity
	secondSellOperation := models.NewSell(secondSellUnitCost, secondSellQuantity)

	// When I apply the buy and sell operations to the capital gains aggregate
	operations := []models.Operation{
		buyOperation,
		firstSellOperation,
		secondSellOperation,
	}
	capitalGains.ApplyOperations(operations)

	// Then the aggregate should register three tax exempted events with zero tax amounts
	taxEvents := capitalGains.Events()
	assert.Len(t, taxEvents, 3)

	taxAmounts := extractTaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.0, // buy
		0.0, // first sell
		0.0, // second sell
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainsApplyOperationsGivenHighProfitThenSubsequentLossWhenApplyOperationsThenOnlyFirstProfitIsTaxed(t *testing.T) {
	t.Parallel()

	// Given I have an empty capital gains aggregate
	capitalGains := models.NewCapitalGains()

	// And a buy unit cost of 10.0
	buyUnitCost := models.NewUnitCost(10.0)

	// And a buy quantity of 10000
	buyQuantity := models.NewQuantity(10000)

	// And I create a buy operation with the buy unit cost and buy quantity
	buyOperation := models.NewBuy(buyUnitCost, buyQuantity)

	// And a first sell unit cost of 20.0
	firstSellUnitCost := models.NewUnitCost(20.0)

	// And a first sell quantity of 5000
	firstSellQuantity := models.NewQuantity(5000)

	// And I create the first sell operation with the first sell unit cost and first sell quantity
	firstSellOperation := models.NewSell(firstSellUnitCost, firstSellQuantity)

	// And a second sell unit cost of 5.0
	secondSellUnitCost := models.NewUnitCost(5.0)

	// And a second sell quantity of 5000
	secondSellQuantity := models.NewQuantity(5000)

	// And I create the second sell operation with the second sell unit cost and second sell quantity
	secondSellOperation := models.NewSell(secondSellUnitCost, secondSellQuantity)

	// When I apply the buy and sell operations to the capital gains aggregate
	operations := []models.Operation{
		buyOperation,
		firstSellOperation,
		secondSellOperation,
	}
	capitalGains.ApplyOperations(operations)

	// Then the aggregate should register one taxed sell and one tax exempted sell with the expected amounts
	taxEvents := capitalGains.Events()
	assert.Len(t, taxEvents, 3)

	taxAmounts := extractTaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.0,     // buy
		10000.0, // first sell
		0.0,     // second sell
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainsApplyOperationsGivenLossThenProfitWhenApplyOperationsThenLossIsDeductedFromProfitBeforeTax(t *testing.T) {
	t.Parallel()

	// Given I have an empty capital gains aggregate
	capitalGains := models.NewCapitalGains()

	// And a buy unit cost of 10.0
	buyUnitCost := models.NewUnitCost(10.0)

	// And a buy quantity of 10000
	buyQuantity := models.NewQuantity(10000)

	// And I create a buy operation with the buy unit cost and buy quantity
	buyOperation := models.NewBuy(buyUnitCost, buyQuantity)

	// And a first sell unit cost of 5.0
	firstSellUnitCost := models.NewUnitCost(5.0)

	// And a first sell quantity of 5000
	firstSellQuantity := models.NewQuantity(5000)

	// And I create the first sell operation with the first sell unit cost and first sell quantity
	firstSellOperation := models.NewSell(firstSellUnitCost, firstSellQuantity)

	// And a second sell unit cost of 20.0
	secondSellUnitCost := models.NewUnitCost(20.0)

	// And a second sell quantity of 3000
	secondSellQuantity := models.NewQuantity(3000)

	// And I create the second sell operation with the second sell unit cost and second sell quantity
	secondSellOperation := models.NewSell(secondSellUnitCost, secondSellQuantity)

	// When I apply the buy and sell operations to the capital gains aggregate
	operations := []models.Operation{
		buyOperation,
		firstSellOperation,
		secondSellOperation,
	}
	capitalGains.ApplyOperations(operations)

	// Then the aggregate should register the expected tax amounts, compensating the loss before taxing the profit
	taxEvents := capitalGains.Events()
	assert.Len(t, taxEvents, 3)

	taxAmounts := extractTaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.0,   // buy
		0.0,   // first sell (loss, no tax)
		1000., // second sell (profit after loss compensation)
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainsApplyOperationsGivenTwoBuysAndSellAtWeightedAverageWhenApplyOperationsThenNoTaxIsCharged(t *testing.T) {
	t.Parallel()

	// Given I have an empty capital gains aggregate
	capitalGains := models.NewCapitalGains()

	// And a first buy unit cost of 10.0
	firstBuyUnitCost := models.NewUnitCost(10.0)

	// And a first buy quantity of 10000
	firstBuyQuantity := models.NewQuantity(10000)

	// And I create the first buy operation with the first buy unit cost and first buy quantity
	firstBuyOperation := models.NewBuy(firstBuyUnitCost, firstBuyQuantity)

	// And a second buy unit cost of 25.0
	secondBuyUnitCost := models.NewUnitCost(25.0)

	// And a second buy quantity of 5000
	secondBuyQuantity := models.NewQuantity(5000)

	// And I create the second buy operation with the second buy unit cost and second buy quantity
	secondBuyOperation := models.NewBuy(secondBuyUnitCost, secondBuyQuantity)

	// And a sell unit cost of 15.0 (the weighted average)
	sellUnitCost := models.NewUnitCost(15.0)

	// And a sell quantity of 15000
	sellQuantity := models.NewQuantity(15000)

	// And I create the sell operation with the sell unit cost and sell quantity
	sellOperation := models.NewSell(sellUnitCost, sellQuantity)

	// When I apply the buy and sell operations to the capital gains aggregate
	operations := []models.Operation{
		firstBuyOperation,
		secondBuyOperation,
		sellOperation,
	}
	capitalGains.ApplyOperations(operations)

	// Then the aggregate should register only tax exempted events
	taxEvents := capitalGains.Events()
	assert.Len(t, taxEvents, 3)

	taxAmounts := extractTaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.0, // first buy
		0.0, // second buy
		0.0, // sell at weighted average
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainsApplyOperationsGivenSellAtAverageThenSellAboveAverageWhenApplyOperationsThenOnlySecondSellIsTaxed(t *testing.T) {
	t.Parallel()

	// Given I have an empty capital gains aggregate
	capitalGains := models.NewCapitalGains()

	// And a first buy unit cost of 10.0
	firstBuyUnitCost := models.NewUnitCost(10.0)

	// And a first buy quantity of 10000
	firstBuyQuantity := models.NewQuantity(10000)

	// And I create the first buy operation with the first buy unit cost and first buy quantity
	firstBuyOperation := models.NewBuy(firstBuyUnitCost, firstBuyQuantity)

	// And a second buy unit cost of 25.0
	secondBuyUnitCost := models.NewUnitCost(25.0)

	// And a second buy quantity of 5000
	secondBuyQuantity := models.NewQuantity(5000)

	// And I create the second buy operation with the second buy unit cost and second buy quantity
	secondBuyOperation := models.NewBuy(secondBuyUnitCost, secondBuyQuantity)

	// And a first sell unit cost of 15.0 (at average)
	firstSellUnitCost := models.NewUnitCost(15.0)

	// And a first sell quantity of 10000
	firstSellQuantity := models.NewQuantity(10000)

	// And I create the first sell operation with the first sell unit cost and first sell quantity
	firstSellOperation := models.NewSell(firstSellUnitCost, firstSellQuantity)

	// And a second sell unit cost of 25.0 (above average)
	secondSellUnitCost := models.NewUnitCost(25.0)

	// And a second sell quantity of 5000
	secondSellQuantity := models.NewQuantity(5000)

	// And I create the second sell operation with the second sell unit cost and second sell quantity
	secondSellOperation := models.NewSell(secondSellUnitCost, secondSellQuantity)

	// When I apply the buy and sell operations to the capital gains aggregate
	operations := []models.Operation{
		firstBuyOperation,
		secondBuyOperation,
		firstSellOperation,
		secondSellOperation,
	}
	capitalGains.ApplyOperations(operations)

	// Then only the second sell should generate tax and the tax amount should match the specification
	taxEvents := capitalGains.Events()
	assert.Len(t, taxEvents, 4)

	taxAmounts := extractTaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.0,     // first buy
		0.0,     // second buy
		0.0,     // first sell at average
		10000.0, // second sell above average
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainsApplyOperationsGivenMultipleCyclesWithLossCompensationAndThresholdWhenApplyOperationsThenTaxesMatchSpecification(t *testing.T) {
	t.Parallel()

	// Given I have an empty capital gains aggregate
	capitalGains := models.NewCapitalGains()

	// And a first cycle buy unit cost of 10.0
	firstCycleBuyUnitCost := models.NewUnitCost(10.0)

	// And a first cycle buy quantity of 10000
	firstCycleBuyQuantity := models.NewQuantity(10000)

	// And I create a first cycle buy operation with the first cycle buy unit cost and first cycle buy quantity
	firstCycleBuyOperation := models.NewBuy(firstCycleBuyUnitCost, firstCycleBuyQuantity)

	// And a first cycle first sell unit cost of 2.0
	firstCycleFirstSellUnitCost := models.NewUnitCost(2.0)

	// And a first cycle first sell quantity of 5000
	firstCycleFirstSellQuantity := models.NewQuantity(5000)

	// And I create a first cycle first sell operation with the first cycle first sell unit cost and first cycle sell quantity
	firstCycleFirstSellOperation := models.NewSell(firstCycleFirstSellUnitCost, firstCycleFirstSellQuantity)

	// And a first cycle second sell unit cost of 20.0
	firstCycleSecondSellUnitCost := models.NewUnitCost(20.0)

	// And a first cycle second sell quantity of 2000
	firstCycleSecondSellQuantity := models.NewQuantity(2000)

	// And I create a first cycle second sell operation with the first cycle second sell unit cost and first cycle second sell quantity
	firstCycleSecondSellOperation := models.NewSell(firstCycleSecondSellUnitCost, firstCycleSecondSellQuantity)

	// And a first cycle third sell unit cost of 20.0
	firstCycleThirdSellUnitCost := models.NewUnitCost(20.0)

	// And a first cycle third sell quantity of 2000
	firstCycleThirdSellQuantity := models.NewQuantity(2000)

	// And I create a first cycle third sell operation with the first cycle third sell unit cost and first cycle third sell quantity
	firstCycleThirdSellOperation := models.NewSell(firstCycleThirdSellUnitCost, firstCycleThirdSellQuantity)

	// And a first cycle fourth sell unit cost of 25.0
	firstCycleFourthSellUnitCost := models.NewUnitCost(25.0)

	// And a first cycle fourth sell quantity of 1000
	firstCycleFourthSellQuantity := models.NewQuantity(1000)

	// And I create a first cycle fourth sell operation with the first cycle fourth sell unit cost and first cycle fourth sell quantity
	firstCycleFourthSellOperation := models.NewSell(firstCycleFourthSellUnitCost, firstCycleFourthSellQuantity)

	// And a second cycle buy unit cost of 20.0
	secondCycleBuyUnitCost := models.NewUnitCost(20.0)

	// And a second cycle buy quantity of 10000
	secondCycleBuyQuantity := models.NewQuantity(10000)

	// And I create a second cycle buy operation with the second cycle buy unit cost and second cycle buy quantity
	secondCycleBuyOperation := models.NewBuy(secondCycleBuyUnitCost, secondCycleBuyQuantity)

	// And a second cycle first sell unit cost of 15.0
	secondCycleFirstSellUnitCost := models.NewUnitCost(15.0)

	// And a second cycle first sell quantity of 5000
	secondCycleFirstSellQuantity := models.NewQuantity(5000)

	// And I create a second cycle first sell operation with the second cycle first sell unit cost and second cycle first sell quantity
	secondCycleFirstSellOperation := models.NewSell(secondCycleFirstSellUnitCost, secondCycleFirstSellQuantity)

	// And a second cycle second sell unit cost of 30.0
	secondCycleSecondSellUnitCost := models.NewUnitCost(30.0)

	// And a second cycle second sell quantity of 4350
	secondCycleSecondSellQuantity := models.NewQuantity(4350)

	// And I create a second cycle second sell operation with the second cycle second sell unit cost and second cycle second sell quantity
	secondCycleSecondSellOperation := models.NewSell(secondCycleSecondSellUnitCost, secondCycleSecondSellQuantity)

	// And a second cycle third sell unit cost of 30.0
	secondCycleThirdSellUnitCost := models.NewUnitCost(30.0)

	// And a second cycle third sell quantity of 650
	secondCycleThirdSellQuantity := models.NewQuantity(650)

	// And I create a second cycle third sell operation with the second cycle third sell unit cost and second cycle third sell quantity
	secondCycleThirdSellOperation := models.NewSell(secondCycleThirdSellUnitCost, secondCycleThirdSellQuantity)

	// When I apply all operations from both cycles to the capital gains aggregate
	operations := []models.Operation{
		firstCycleBuyOperation,
		firstCycleFirstSellOperation,
		firstCycleSecondSellOperation,
		firstCycleThirdSellOperation,
		firstCycleFourthSellOperation,
		secondCycleBuyOperation,
		secondCycleFirstSellOperation,
		secondCycleSecondSellOperation,
		secondCycleThirdSellOperation,
	}
	capitalGains.ApplyOperations(operations)

	// Then the aggregate tax amounts should match the specification
	taxEvents := capitalGains.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := extractTaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.0,    // first cycle buy
		0.0,    // first cycle sell (loss)
		0.0,    // first cycle second sell (loss absorbed by accumulated loss)
		0.0,    // first cycle third sell (loss absorbed)
		3000.0, // first cycle fourth sell (taxed profit)
		0.0,    // second cycle buy
		0.0,    // second cycle first sell (under threshold or compensated)
		3700.0, // second cycle second sell (taxed profit)
		0.0,    // second cycle third sell (no tax)
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainsApplyOperationsGivenTwoIndependentHighProfitCyclesWhenApplyOperationsThenTaxIsCalculatedForEachCycle(t *testing.T) {
	t.Parallel()

	// Given I have an empty capital gains aggregate
	capitalGains := models.NewCapitalGains()

	// And a first cycle buy unit cost of 10.0
	firstCycleBuyUnitCost := models.NewUnitCost(10.0)

	// And a first cycle buy quantity of 10000
	firstCycleBuyQuantity := models.NewQuantity(10000)

	// And I create a first cycle buy operation with the first cycle buy unit cost and first cycle buy quantity
	firstCycleBuyOperation := models.NewBuy(firstCycleBuyUnitCost, firstCycleBuyQuantity)

	// And a first cycle sell unit cost of 50.0
	firstCycleSellUnitCost := models.NewUnitCost(50.0)

	// And a first cycle sell quantity of 10000
	firstCycleSellQuantity := models.NewQuantity(10000)

	// And I create a first cycle sell operation with the first cycle sell unit cost and first cycle sell quantity
	firstCycleSellOperation := models.NewSell(firstCycleSellUnitCost, firstCycleSellQuantity)

	// And a second cycle buy unit cost of 20.0
	secondCycleBuyUnitCost := models.NewUnitCost(20.0)

	// And a second cycle buy quantity of 10000
	secondCycleBuyQuantity := models.NewQuantity(10000)

	// And I create a second cycle buy operation with the second cycle buy unit cost and second cycle buy quantity
	secondCycleBuyOperation := models.NewBuy(secondCycleBuyUnitCost, secondCycleBuyQuantity)

	// And a second cycle sell unit cost of 50.0
	secondCycleSellUnitCost := models.NewUnitCost(50.0)

	// And a second cycle sell quantity of 10000
	secondCycleSellQuantity := models.NewQuantity(10000)

	// And I create a second cycle sell operation with the second cycle sell unit cost and second cycle sell quantity
	secondCycleSellOperation := models.NewSell(secondCycleSellUnitCost, secondCycleSellQuantity)

	// When I apply the operations from both cycles to the capital gains aggregate
	operations := []models.Operation{
		firstCycleBuyOperation,
		firstCycleSellOperation,
		secondCycleBuyOperation,
		secondCycleSellOperation,
	}
	capitalGains.ApplyOperations(operations)

	// Then each high profit cycle should generate the expected tax amount
	taxEvents := capitalGains.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := extractTaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.0,    // first cycle buy
		80000., // first cycle sell
		0.0,    // second cycle buy
		60000., // second cycle sell
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func extractTaxAmountsFromEvents(taxEvents []events.Event) []float64 {
	taxAmounts := make([]float64, len(taxEvents))

	for index, taxEvent := range taxEvents {
		taxAmounts[index] = taxEvent.Amount()
	}

	return taxAmounts
}
