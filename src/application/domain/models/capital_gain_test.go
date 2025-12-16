package models_test

import (
	"capital-gains/test"
	"testing"

	"capital-gains/src/application/domain/models"

	"github.com/stretchr/testify/assert"
)

func TestCapitalGainApplyOperationsGivenProfitsWithProceedsBelowThresholdWhenApplyOperationsThenNoTaxIsCharged(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a buy quantity of 100
	buyQuantity := models.NewQuantity(100)

	// And a buy unit cost of 10.00
	buyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a buy operation with the buy quantity and buy unit cost
	buyOperation := models.NewBuy(buyQuantity, buyUnitCost)

	// And a first sell quantity of 50
	firstSellQuantity := models.NewQuantity(50)

	// And a first sell unit cost of 15.00
	firstSellUnitCost := models.NewMonetaryValue(15.00)

	// And I create the first sell operation with the first sell quantity and first sell unit cost
	firstSellOperation := models.NewSell(firstSellQuantity, firstSellUnitCost)

	// And a second sell quantity of 50
	secondSellQuantity := models.NewQuantity(50)

	// And a second sell unit cost of 15.00
	secondSellUnitCost := models.NewMonetaryValue(15.00)

	// And I create the second sell operation with the second sell quantity and second sell unit cost
	secondSellOperation := models.NewSell(secondSellQuantity, secondSellUnitCost)

	// When I apply the buy and sell operations to the capital gain
	operations := []models.Operation{
		buyOperation,
		firstSellOperation,
		secondSellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then the tax result should have three exempt operations with zero tax amounts
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, 3)

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00, // buy
		0.00, // first sell
		0.00, // second sell
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenHighProfitThenSubsequentLossWhenApplyOperationsThenOnlyFirstProfitIsTaxed(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a buy quantity of 10000
	buyQuantity := models.NewQuantity(10000)

	// And a buy unit cost of 10.00
	buyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a buy operation with the buy quantity and buy unit cost
	buyOperation := models.NewBuy(buyQuantity, buyUnitCost)

	// And a first sell quantity of 5000
	firstSellQuantity := models.NewQuantity(5000)

	// And a first sell unit cost of 20.00
	firstSellUnitCost := models.NewMonetaryValue(20.00)

	// And I create the first sell operation with the first sell quantity and first sell unit cost
	firstSellOperation := models.NewSell(firstSellQuantity, firstSellUnitCost)

	// And a second sell quantity of 5000
	secondSellQuantity := models.NewQuantity(5000)

	// And a second sell unit cost of 5.00
	secondSellUnitCost := models.NewMonetaryValue(5.00)

	// And I create the second sell operation with the second sell quantity and second sell unit cost
	secondSellOperation := models.NewSell(secondSellQuantity, secondSellUnitCost)

	// When I apply the buy and sell operations to the capital gain
	operations := []models.Operation{
		buyOperation,
		firstSellOperation,
		secondSellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then the tax result should have one taxed sell and one exempt sell with the expected amounts
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, 3)

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00,     // buy
		10000.00, // first sell
		0.00,     // second sell
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenLossThenProfitWhenApplyOperationsThenLossIsDeductedFromProfitBeforeTax(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a buy quantity of 10000
	buyQuantity := models.NewQuantity(10000)

	// And a buy unit cost of 10.00
	buyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a buy operation with the buy quantity and buy unit cost
	buyOperation := models.NewBuy(buyQuantity, buyUnitCost)

	// And a first sell quantity of 5000
	firstSellQuantity := models.NewQuantity(5000)

	// And a first sell unit cost of 5.00
	firstSellUnitCost := models.NewMonetaryValue(5.00)

	// And I create the first sell operation with the first sell quantity and first sell unit cost
	firstSellOperation := models.NewSell(firstSellQuantity, firstSellUnitCost)

	// And a second sell quantity of 3000
	secondSellQuantity := models.NewQuantity(3000)

	// And a second sell unit cost of 20.00
	secondSellUnitCost := models.NewMonetaryValue(20.00)

	// And I create the second sell operation with the second sell quantity and second sell unit cost
	secondSellOperation := models.NewSell(secondSellQuantity, secondSellUnitCost)

	// When I apply the buy and sell operations to the capital gain
	operations := []models.Operation{
		buyOperation,
		firstSellOperation,
		secondSellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then the tax result should match the expected loss compensation before taxing the profit
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, 3)

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00,    // buy
		0.00,    // first sell (loss, no tax)
		1000.00, // second sell (profit after loss compensation)
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenTwoBuysAndSellAtWeightedAverageWhenApplyOperationsThenNoTaxIsCharged(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a first buy quantity of 10000
	firstBuyQuantity := models.NewQuantity(10000)

	// And a first buy unit cost of 10.00
	firstBuyUnitCost := models.NewMonetaryValue(10.00)

	// And I create the first buy operation with the first buy quantity and first buy unit cost
	firstBuyOperation := models.NewBuy(firstBuyQuantity, firstBuyUnitCost)

	// And a second buy quantity of 5000
	secondBuyQuantity := models.NewQuantity(5000)

	// And a second buy unit cost of 25.00
	secondBuyUnitCost := models.NewMonetaryValue(25.00)

	// And I create the second buy operation with the second buy quantity and second buy unit cost
	secondBuyOperation := models.NewBuy(secondBuyQuantity, secondBuyUnitCost)

	// And a sell quantity of 15000
	sellQuantity := models.NewQuantity(15000)

	// And a sell unit cost of 15.00 (the weighted average)
	sellUnitCost := models.NewMonetaryValue(15.00)

	// And I create the sell operation with the sell quantity and sell unit cost
	sellOperation := models.NewSell(sellQuantity, sellUnitCost)

	// When I apply the buy and sell operations to the capital gain
	operations := []models.Operation{
		firstBuyOperation,
		secondBuyOperation,
		sellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then the tax result should have only exempt operations
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, 3)

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00, // first buy
		0.00, // second buy
		0.00, // sell at weighted average
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenSellAtAverageThenSellAboveAverageWhenApplyOperationsThenOnlySecondSellIsTaxed(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a first buy quantity of 10000
	firstBuyQuantity := models.NewQuantity(10000)

	// And a first buy unit cost of 10.00
	firstBuyUnitCost := models.NewMonetaryValue(10.00)

	// And I create the first buy operation with the first buy quantity and first buy unit cost
	firstBuyOperation := models.NewBuy(firstBuyQuantity, firstBuyUnitCost)

	// And a second buy quantity of 5000
	secondBuyQuantity := models.NewQuantity(5000)

	// And a second buy unit cost of 25.00
	secondBuyUnitCost := models.NewMonetaryValue(25.00)

	// And I create the second buy operation with the second buy quantity and second buy unit cost
	secondBuyOperation := models.NewBuy(secondBuyQuantity, secondBuyUnitCost)

	// And a first sell quantity of 10000
	firstSellQuantity := models.NewQuantity(10000)

	// And a first sell unit cost of 15.00 (at average)
	firstSellUnitCost := models.NewMonetaryValue(15.00)

	// And I create the first sell operation with the first sell quantity and first sell unit cost
	firstSellOperation := models.NewSell(firstSellQuantity, firstSellUnitCost)

	// And a second sell quantity of 5000
	secondSellQuantity := models.NewQuantity(5000)

	// And a second sell unit cost of 25.00 (above average)
	secondSellUnitCost := models.NewMonetaryValue(25.00)

	// And I create the second sell operation with the second sell quantity and second sell unit cost
	secondSellOperation := models.NewSell(secondSellQuantity, secondSellUnitCost)

	// When I apply the buy and sell operations to the capital gain
	operations := []models.Operation{
		firstBuyOperation,
		secondBuyOperation,
		firstSellOperation,
		secondSellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then only the second sell should generate tax and the tax amount should match the specification
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, 4)

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00,     // first buy
		0.00,     // second buy
		0.00,     // first sell at average
		10000.00, // second sell above average
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenMultipleCyclesWithLossCompensationAndThresholdWhenApplyOperationsThenTaxesMatchSpecification(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a first cycle buy quantity of 10000
	firstCycleBuyQuantity := models.NewQuantity(10000)

	// And a first cycle buy unit cost of 10.00
	firstCycleBuyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a first cycle buy operation with the first cycle buy quantity and first cycle buy unit cost
	firstCycleBuyOperation := models.NewBuy(firstCycleBuyQuantity, firstCycleBuyUnitCost)

	// And a first cycle first sell quantity of 5000
	firstCycleFirstSellQuantity := models.NewQuantity(5000)

	// And a first cycle first sell unit cost of 2.00
	firstCycleFirstSellUnitCost := models.NewMonetaryValue(2.00)

	// And I create a first cycle first sell operation with the first cycle first sell quantity and first cycle sell unit cost
	firstCycleFirstSellOperation := models.NewSell(firstCycleFirstSellQuantity, firstCycleFirstSellUnitCost)

	// And a first cycle second sell quantity of 2000
	firstCycleSecondSellQuantity := models.NewQuantity(2000)

	// And a first cycle second sell unit cost of 20.00
	firstCycleSecondSellUnitCost := models.NewMonetaryValue(20.00)

	// And I create a first cycle second sell operation with the first cycle second sell quantity and first cycle second sell unit cost
	firstCycleSecondSellOperation := models.NewSell(firstCycleSecondSellQuantity, firstCycleSecondSellUnitCost)

	// And a first cycle third sell quantity of 2000
	firstCycleThirdSellQuantity := models.NewQuantity(2000)

	// And a first cycle third sell unit cost of 20.00
	firstCycleThirdSellUnitCost := models.NewMonetaryValue(20.00)

	// And I create a first cycle third sell operation with the first cycle third sell quantity and first cycle third sell unit cost
	firstCycleThirdSellOperation := models.NewSell(firstCycleThirdSellQuantity, firstCycleThirdSellUnitCost)

	// And a first cycle fourth sell quantity of 1000
	firstCycleFourthSellQuantity := models.NewQuantity(1000)

	// And a first cycle fourth sell unit cost of 25.00
	firstCycleFourthSellUnitCost := models.NewMonetaryValue(25.00)

	// And I create a first cycle fourth sell operation with the first cycle fourth sell quantity and first cycle fourth sell unit cost
	firstCycleFourthSellOperation := models.NewSell(firstCycleFourthSellQuantity, firstCycleFourthSellUnitCost)

	// And a second cycle buy quantity of 10000
	secondCycleBuyQuantity := models.NewQuantity(10000)

	// And a second cycle buy unit cost of 20.00
	secondCycleBuyUnitCost := models.NewMonetaryValue(20.00)

	// And I create a second cycle buy operation with the second cycle buy quantity and second cycle buy unit cost
	secondCycleBuyOperation := models.NewBuy(secondCycleBuyQuantity, secondCycleBuyUnitCost)

	// And a second cycle first sell quantity of 5000
	secondCycleFirstSellQuantity := models.NewQuantity(5000)

	// And a second cycle first sell unit cost of 15.00
	secondCycleFirstSellUnitCost := models.NewMonetaryValue(15.00)

	// And I create a second cycle first sell operation with the second cycle first sell quantity and second cycle first sell unit cost
	secondCycleFirstSellOperation := models.NewSell(secondCycleFirstSellQuantity, secondCycleFirstSellUnitCost)

	// And a second cycle second sell quantity of 4350
	secondCycleSecondSellQuantity := models.NewQuantity(4350)

	// And a second cycle second sell unit cost of 30.0
	secondCycleSecondSellUnitCost := models.NewMonetaryValue(30.0)

	// And I create a second cycle second sell operation with the second cycle second sell quantity and second cycle second sell unit cost
	secondCycleSecondSellOperation := models.NewSell(secondCycleSecondSellQuantity, secondCycleSecondSellUnitCost)

	// And a second cycle third sell quantity of 650
	secondCycleThirdSellQuantity := models.NewQuantity(650)

	// And a second cycle third sell unit cost of 30.0
	secondCycleThirdSellUnitCost := models.NewMonetaryValue(30.0)

	// And I create a second cycle third sell operation with the second cycle third sell quantity and second cycle third sell unit cost
	secondCycleThirdSellOperation := models.NewSell(secondCycleThirdSellQuantity, secondCycleThirdSellUnitCost)

	// When I apply all operations from both cycles to the capital gain
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
	capitalGain.ApplyOperations(operations)

	// Then the tax amounts should match the expected values for both cycles
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00,    // first cycle buy
		0.00,    // first cycle sell (loss)
		0.00,    // first cycle second sell (loss absorbed by accumulated loss)
		0.00,    // first cycle third sell (loss absorbed)
		3000.00, // first cycle fourth sell (taxed profit)
		0.00,    // second cycle buy
		0.00,    // second cycle first sell (under threshold or compensated)
		3700.00, // second cycle second sell (taxed profit)
		0.00,    // second cycle third sell (no tax)
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenTwoIndependentHighProfitCyclesWhenApplyOperationsThenTaxIsCalculatedForEachCycle(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a first cycle buy quantity of 10000
	firstCycleBuyQuantity := models.NewQuantity(10000)

	// And a first cycle buy unit cost of 10.00
	firstCycleBuyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a first cycle buy operation with the first cycle buy quantity and first cycle buy unit cost
	firstCycleBuyOperation := models.NewBuy(firstCycleBuyQuantity, firstCycleBuyUnitCost)

	// And a first cycle sell quantity of 10000
	firstCycleSellQuantity := models.NewQuantity(10000)

	// And a first cycle sell unit cost of 50.00
	firstCycleSellUnitCost := models.NewMonetaryValue(50.00)

	// And I create a first cycle sell operation with the first cycle sell quantity and first cycle sell unit cost
	firstCycleSellOperation := models.NewSell(firstCycleSellQuantity, firstCycleSellUnitCost)

	// And a second cycle buy quantity of 10000
	secondCycleBuyQuantity := models.NewQuantity(10000)

	// And a second cycle buy unit cost of 20.00
	secondCycleBuyUnitCost := models.NewMonetaryValue(20.00)

	// And I create a second cycle buy operation with the second cycle buy quantity and second cycle buy unit cost
	secondCycleBuyOperation := models.NewBuy(secondCycleBuyQuantity, secondCycleBuyUnitCost)

	// And a second cycle sell quantity of 10000
	secondCycleSellQuantity := models.NewQuantity(10000)

	// And a second cycle sell unit cost of 50.00
	secondCycleSellUnitCost := models.NewMonetaryValue(50.00)

	// And I create a second cycle sell operation with the second cycle sell quantity and second cycle sell unit cost
	secondCycleSellOperation := models.NewSell(secondCycleSellQuantity, secondCycleSellUnitCost)

	// When I apply the operations from both cycles to the capital gain
	operations := []models.Operation{
		firstCycleBuyOperation,
		firstCycleSellOperation,
		secondCycleBuyOperation,
		secondCycleSellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then each high profit cycle should generate the expected tax amount
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00,     // first cycle buy
		80000.00, // first cycle sell
		0.00,     // second cycle buy
		60000.00, // second cycle sell
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenZeroQuantityBuyWhenApplyOperationsThenTaxesAreUnaffected(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a zero quantity buy quantity of 0
	zeroQuantityBuyQuantity := models.NewQuantity(0)

	// And a zero quantity buy unit cost of 10.00
	zeroQuantityBuyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a zero quantity buy operation with the zero quantity buy quantity and zero quantity buy unit cost
	zeroQuantityBuyOperation := models.NewBuy(zeroQuantityBuyQuantity, zeroQuantityBuyUnitCost)

	// And a regular buy quantity of 10000
	regularBuyQuantity := models.NewQuantity(10000)

	// And a regular buy unit cost of 10.00
	regularBuyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a regular buy operation with the regular buy quantity and regular buy unit cost
	regularBuyOperation := models.NewBuy(regularBuyQuantity, regularBuyUnitCost)

	// And a sell quantity of 5000
	sellQuantity := models.NewQuantity(5000)

	// And a sell unit cost of 20.00
	sellUnitCost := models.NewMonetaryValue(20.00)

	// And I create a sell operation with the sell quantity and sell unit cost
	sellOperation := models.NewSell(sellQuantity, sellUnitCost)

	// When I apply the zero quantity buy, the regular buy and the sell operations to the capital gain
	operations := []models.Operation{
		zeroQuantityBuyOperation,
		regularBuyOperation,
		sellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then the tax result should be the same as if the zero quantity buy did not exist
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00,     // zero quantity buy does not change the tax
		0.00,     // regular buy
		10000.00, // sell above average with proceeds above threshold
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}
