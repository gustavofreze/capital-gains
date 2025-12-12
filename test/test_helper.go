package test

import "capital-gains/src/application/domain/events"

// TaxAmountsFromEvents extracts the tax amounts from a slice of tax events.
func TaxAmountsFromEvents(taxEvents []events.Event) []float64 {
	taxAmounts := make([]float64, len(taxEvents))

	for index, taxEvent := range taxEvents {
		taxAmounts[index] = taxEvent.Amount()
	}

	return taxAmounts
}
