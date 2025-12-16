package models_test

import (
	"testing"

	"capital-gains/test"

	"capital-gains/src/application/domain/models"

	"github.com/stretchr/testify/assert"
)

func TestCapitalGainApplyOperationsGivenTaxFreeProfitSalesBelowThresholdWhenApplyOperationsThenNoTaxIsCharged(t *testing.T) {
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

func TestCapitalGainApplyOperationsGivenTaxableProfitSaleThenSubsequentLossWhenApplyOperationsThenOnlyFirstSaleIsTaxed(t *testing.T) {
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

func TestCapitalGainApplyOperationsGivenLossCarryforwardThenTaxableProfitSaleWhenApplyOperationsThenLossIsOffsetBeforeTax(t *testing.T) {
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

func TestCapitalGainApplyOperationsGivenWeightedAverageCostThenSellAtAverageWhenApplyOperationsThenNoTaxIsCharged(t *testing.T) {
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

func TestCapitalGainApplyOperationsGivenSellAtAverageThenTaxableProfitSaleWhenApplyOperationsThenOnlySecondSaleIsTaxed(t *testing.T) {
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

func TestCapitalGainApplyOperationsGivenOfficialSampleScenariosWhenApplyOperationsThenTaxesMatchSpecification(t *testing.T) {
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

func TestCapitalGainApplyOperationsGivenTwoTaxableProfitCyclesWhenApplyOperationsThenTaxIsCalculatedForEachCycle(t *testing.T) {
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

func TestCapitalGainApplyOperationsGivenZeroQuantityBuyOperationWhenApplyOperationsThenTaxesAreUnaffected(t *testing.T) {
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

func TestCapitalGainApplyOperationsGivenAccumulatedLossThenTaxFreeProfitSaleWhenApplyOperationsThenLossIsPreservedForFutureTaxableSale(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a buy quantity of 1000
	buyQuantity := models.NewQuantity(1000)

	// And a buy unit cost of 10.00
	buyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a buy operation with the buy quantity and buy unit cost
	firstBuyOperation := models.NewBuy(buyQuantity, buyUnitCost)

	// And a first sell quantity of 500
	firstSellQuantity := models.NewQuantity(500)

	// And a first sell unit cost of 5.00 (loss of 2500.00)
	firstSellUnitCost := models.NewMonetaryValue(5.00)

	// And I create the first sell operation with the first sell quantity and first sell unit cost
	firstSellOperation := models.NewSell(firstSellQuantity, firstSellUnitCost)

	// And a second sell quantity of 500
	secondSellQuantity := models.NewQuantity(500)

	// And a second sell unit cost of 15.00 (profit of 2500.00, but the proceeds are 7500.00, below the exemption threshold)
	secondSellUnitCost := models.NewMonetaryValue(15.00)

	// And I create the second sell operation with the second sell quantity and second sell unit cost
	secondSellOperation := models.NewSell(secondSellQuantity, secondSellUnitCost)

	// And a second buy quantity of 1000
	secondBuyQuantity := models.NewQuantity(1000)

	// And a second buy unit cost of 10.00
	secondBuyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a second buy operation with the second buy quantity and second buy unit cost
	secondBuyOperation := models.NewBuy(secondBuyQuantity, secondBuyUnitCost)

	// And a third sell quantity of 1000
	thirdSellQuantity := models.NewQuantity(1000)

	// And a third sell unit cost of 30.00 (profit of 20000.00, proceeds above threshold)
	thirdSellUnitCost := models.NewMonetaryValue(30.00)

	// And I create the third sell operation with the third sell quantity and third sell unit cost
	thirdSellOperation := models.NewSell(thirdSellQuantity, thirdSellUnitCost)

	// When I apply all operations
	operations := []models.Operation{
		firstBuyOperation,
		firstSellOperation,
		secondSellOperation,
		secondBuyOperation,
		thirdSellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then the exempt profit must not reduce the accumulated loss.
	// So the taxable profit in the last sell is 20000.00 - 2500.00 = 17500.00,
	// and the tax is 20% of 17500.00 = 3500.00.
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00,    // first buy
		0.00,    // first sell (loss)
		0.00,    // second sell (profit, but exempt by threshold)
		0.00,    // second buy
		3500.00, // third sell (taxed, with loss carryforward preserved)
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenSaleProceedsEqualToThresholdWhenApplyOperationsThenNoTaxIsCharged(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a buy quantity of 1000
	buyQuantity := models.NewQuantity(1000)

	// And a buy unit cost of 10.00
	buyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a buy operation with the buy quantity and buy unit cost
	buyOperation := models.NewBuy(buyQuantity, buyUnitCost)

	// And a sell quantity of 1000
	sellQuantity := models.NewQuantity(1000)

	// And a sell unit cost of 20.00 (proceeds are exactly 20000.00, which is tax-free)
	sellUnitCost := models.NewMonetaryValue(20.00)

	// And I create a sell operation with the sell quantity and sell unit cost
	sellOperation := models.NewSell(sellQuantity, sellUnitCost)

	// When I apply the buy and sell operations to the capital gain
	operations := []models.Operation{
		buyOperation,
		sellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then no tax is charged because the sale proceeds are exactly the threshold
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00, // buy
		0.00, // sell (tax-free by threshold)
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenTaxFreeLossSaleWhenApplyOperationsThenLossIsCarriedForwardToFutureTaxableProfit(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a buy quantity of 1000
	buyQuantity := models.NewQuantity(1000)

	// And a buy unit cost of 10.00
	buyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a buy operation with the buy quantity and buy unit cost
	firstBuyOperation := models.NewBuy(buyQuantity, buyUnitCost)

	// And a sell quantity of 1000
	firstSellQuantity := models.NewQuantity(1000)

	// And a sell unit cost of 5.00 (loss of 5000.00, proceeds below threshold)
	firstSellUnitCost := models.NewMonetaryValue(5.00)

	// And I create a sell operation with the sell quantity and sell unit cost
	firstSellOperation := models.NewSell(firstSellQuantity, firstSellUnitCost)

	// And a second buy quantity of 1000
	secondBuyQuantity := models.NewQuantity(1000)

	// And a second buy unit cost of 10.00
	secondBuyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a second buy operation with the second buy quantity and second buy unit cost
	secondBuyOperation := models.NewBuy(secondBuyQuantity, secondBuyUnitCost)

	// And a second sell quantity of 1000
	secondSellQuantity := models.NewQuantity(1000)

	// And a second sell unit cost of 40.00 (profit of 30000.00, proceeds above threshold)
	secondSellUnitCost := models.NewMonetaryValue(40.00)

	// And I create the second sell operation with the second sell quantity and second sell unit cost
	secondSellOperation := models.NewSell(secondSellQuantity, secondSellUnitCost)

	// When I apply all operations to the capital gain
	operations := []models.Operation{
		firstBuyOperation,
		firstSellOperation,
		secondBuyOperation,
		secondSellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then the tax on the last sell must reflect the carried forward loss:
	// net profit is 30000.00 - 5000.00 = 25000.00 => tax 20% = 5000.00
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00,    // first buy
		0.00,    // first sell (loss, tax-free by threshold)
		0.00,    // second buy
		5000.00, // second sell (taxed after loss offset)
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenAccumulatedLossThenMultipleTaxFreeProfitSalesWhenApplyOperationsThenLossIsPreservedForFutureTaxableSale(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a buy quantity of 1000
	buyQuantity := models.NewQuantity(1000)

	// And a buy unit cost of 10.00
	buyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a buy operation with the buy quantity and buy unit cost
	firstBuyOperation := models.NewBuy(buyQuantity, buyUnitCost)

	// And a first sell quantity of 500
	firstSellQuantity := models.NewQuantity(500)

	// And a first sell unit cost of 5.00 (loss of 2500.00)
	firstSellUnitCost := models.NewMonetaryValue(5.00)

	// And I create the first sell operation with the first sell quantity and first sell unit cost
	firstSellOperation := models.NewSell(firstSellQuantity, firstSellUnitCost)

	// And a second sell quantity of 100
	secondSellQuantity := models.NewQuantity(100)

	// And a second sell unit cost of 15.00 (profit of 500.00, proceeds below threshold)
	secondSellUnitCost := models.NewMonetaryValue(15.00)

	// And I create the second sell operation with the second sell quantity and second sell unit cost
	secondSellOperation := models.NewSell(secondSellQuantity, secondSellUnitCost)

	// And a third sell quantity of 400
	thirdSellQuantity := models.NewQuantity(400)

	// And a third sell unit cost of 15.00 (profit of 2000.00, proceeds below threshold)
	thirdSellUnitCost := models.NewMonetaryValue(15.00)

	// And I create the third sell operation with the third sell quantity and third sell unit cost
	thirdSellOperation := models.NewSell(thirdSellQuantity, thirdSellUnitCost)

	// And a second buy quantity of 1000
	secondBuyQuantity := models.NewQuantity(1000)

	// And a second buy unit cost of 10.00
	secondBuyUnitCost := models.NewMonetaryValue(10.00)

	// And I create a second buy operation with the second buy quantity and second buy unit cost
	secondBuyOperation := models.NewBuy(secondBuyQuantity, secondBuyUnitCost)

	// And a fourth sell quantity of 1000
	fourthSellQuantity := models.NewQuantity(1000)

	// And a fourth sell unit cost of 30.00 (profit of 20000.00, proceeds above threshold)
	fourthSellUnitCost := models.NewMonetaryValue(30.00)

	// And I create the fourth sell operation with the fourth sell quantity and fourth sell unit cost
	fourthSellOperation := models.NewSell(fourthSellQuantity, fourthSellUnitCost)

	// When I apply all operations
	operations := []models.Operation{
		firstBuyOperation,
		firstSellOperation,
		secondSellOperation,
		thirdSellOperation,
		secondBuyOperation,
		fourthSellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then the tax in the last sell must reflect the original loss (2500.00),
	// because tax-free profit sells must not consume accumulated losses:
	// net profit is 20000.00 - 2500.00 = 17500.00 => tax 20% = 3500.00
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00,    // first buy
		0.00,    // first sell (loss)
		0.00,    // second sell (profit, tax-free by threshold)
		0.00,    // third sell (profit, tax-free by threshold)
		0.00,    // second buy
		3500.00, // fourth sell (taxed, with loss carryforward preserved)
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenTaxableSaleWithZeroGainWhenApplyOperationsThenNoTaxIsCharged(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a buy quantity of 2000
	buyQuantity := models.NewQuantity(2000)

	// And a buy unit cost of 20.00
	buyUnitCost := models.NewMonetaryValue(20.00)

	// And I create a buy operation with the buy quantity and buy unit cost
	buyOperation := models.NewBuy(buyQuantity, buyUnitCost)

	// And a sell quantity of 1500
	sellQuantity := models.NewQuantity(1500)

	// And a sell unit cost of 20.00 (sell at weighted average, proceeds above threshold)
	sellUnitCost := models.NewMonetaryValue(20.00)

	// And I create a sell operation with the sell quantity and sell unit cost
	sellOperation := models.NewSell(sellQuantity, sellUnitCost)

	// When I apply the buy and sell operations
	operations := []models.Operation{
		buyOperation,
		sellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then no tax is charged because the gain is zero
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00, // buy
		0.00, // sell at average
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenTaxableLossSaleWhenApplyOperationsThenLossIsAccumulatedAndOffsetsFutureTaxableProfit(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a buy quantity of 10000
	buyQuantity := models.NewQuantity(10000)

	// And a buy unit cost of 5.00
	buyUnitCost := models.NewMonetaryValue(5.00)

	// And I create a buy operation with the buy quantity and buy unit cost
	buyOperation := models.NewBuy(buyQuantity, buyUnitCost)

	// And a first sell quantity of 6000
	firstSellQuantity := models.NewQuantity(6000)

	// And a first sell unit cost of 4.00 (loss of 6000.00, proceeds above threshold)
	firstSellUnitCost := models.NewMonetaryValue(4.00)

	// And I create the first sell operation with the first sell quantity and first sell unit cost
	firstSellOperation := models.NewSell(firstSellQuantity, firstSellUnitCost)

	// And a second sell quantity of 4000
	secondSellQuantity := models.NewQuantity(4000)

	// And a second sell unit cost of 10.00 (profit of 20000.00, proceeds above threshold)
	secondSellUnitCost := models.NewMonetaryValue(10.00)

	// And I create the second sell operation with the second sell quantity and second sell unit cost
	secondSellOperation := models.NewSell(secondSellQuantity, secondSellUnitCost)

	// When I apply the buy and sell operations
	operations := []models.Operation{
		buyOperation,
		firstSellOperation,
		secondSellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then the tax in the second sell must reflect the loss offset:
	// net profit is 20000.00 - 6000.00 = 14000.00 => tax 20% = 2800.00
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00,    // buy
		0.00,    // first sell (loss)
		2800.00, // second sell (taxed after loss offset)
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenLargeAccumulatedLossThenMultipleTaxableProfitSalesWhenApplyOperationsThenLossIsOffsetAcrossSalesUntilExhausted(t *testing.T) {
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

	// And a first sell unit cost of 5.00 (loss of 25000.00, proceeds above threshold)
	firstSellUnitCost := models.NewMonetaryValue(5.00)

	// And I create the first sell operation with the first sell quantity and first sell unit cost
	firstSellOperation := models.NewSell(firstSellQuantity, firstSellUnitCost)

	// And a second sell quantity of 2000
	secondSellQuantity := models.NewQuantity(2000)

	// And a second sell unit cost of 20.00 (profit of 20000.00, proceeds above threshold)
	secondSellUnitCost := models.NewMonetaryValue(20.00)

	// And I create the second sell operation with the second sell quantity and second sell unit cost
	secondSellOperation := models.NewSell(secondSellQuantity, secondSellUnitCost)

	// And a third sell quantity of 1000
	thirdSellQuantity := models.NewQuantity(1000)

	// And a third sell unit cost of 25.00 (profit of 15000.00, proceeds above threshold)
	thirdSellUnitCost := models.NewMonetaryValue(25.00)

	// And I create the third sell operation with the third sell quantity and third sell unit cost
	thirdSellOperation := models.NewSell(thirdSellQuantity, thirdSellUnitCost)

	// When I apply all operations
	operations := []models.Operation{
		buyOperation,
		firstSellOperation,
		secondSellOperation,
		thirdSellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then the accumulated loss must be consumed across multiple taxable profit sells until exhausted:
	// - second sell profit (20000.00) is fully offset by loss (25000.00) => tax 0.00, remaining loss 5000.00
	// - third sell profit (15000.00) is partially offset by remaining loss (5000.00) => net 10000.00 => tax 2000.00
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00,    // buy
		0.00,    // first sell (loss)
		0.00,    // second sell (profit fully offset by loss)
		2000.00, // third sell (taxed after remaining loss offset)
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}

func TestCapitalGainApplyOperationsGivenWeightedAverageCostWithRepeatingDecimalWhenApplyOperationsThenTaxUsesRoundedAverageToTwoDecimals(t *testing.T) {
	t.Parallel()

	// Given I start a new capital gain calculation
	capitalGain := models.NewCapitalGain()

	// And a first buy quantity of 1000
	firstBuyQuantity := models.NewQuantity(1000)

	// And a first buy unit cost of 10.00
	firstBuyUnitCost := models.NewMonetaryValue(10.00)

	// And I create the first buy operation with the first buy quantity and first buy unit cost
	firstBuyOperation := models.NewBuy(firstBuyQuantity, firstBuyUnitCost)

	// And a second buy quantity of 2000
	secondBuyQuantity := models.NewQuantity(2000)

	// And a second buy unit cost of 20.00
	secondBuyUnitCost := models.NewMonetaryValue(20.00)

	// And I create the second buy operation with the second buy quantity and second buy unit cost
	secondBuyOperation := models.NewBuy(secondBuyQuantity, secondBuyUnitCost)

	// And a sell quantity of 3000
	sellQuantity := models.NewQuantity(3000)

	// And a sell unit cost of 30.00 (proceeds above threshold)
	sellUnitCost := models.NewMonetaryValue(30.00)

	// And I create the sell operation with the sell quantity and sell unit cost
	sellOperation := models.NewSell(sellQuantity, sellUnitCost)

	// When I apply the operations
	operations := []models.Operation{
		firstBuyOperation,
		secondBuyOperation,
		sellOperation,
	}
	capitalGain.ApplyOperations(operations)

	// Then the tax must be calculated using the weighted average rounded to two decimal places:
	// wac is (1000*10 + 2000*20) / 3000 = 16.666... => 16.67
	// profit is (30.00 - 16.67) * 3000 = 39990.00 => tax 20% = 7998.00
	taxEvents := capitalGain.Events()
	assert.Len(t, taxEvents, len(operations))

	taxAmounts := test.TaxAmountsFromEvents(taxEvents)
	expectedTaxAmounts := []float64{
		0.00,    // first buy
		0.00,    // second buy
		7998.00, // sell taxed with rounded weighted average
	}

	assert.Equal(t, expectedTaxAmounts, taxAmounts)
}
