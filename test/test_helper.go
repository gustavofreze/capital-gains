package test

import (
	"capital-gains/src/application/domain/events"
	"encoding/json"
)

// TaxAmountsFromEvents extracts the tax amounts from a slice of tax events.
func TaxAmountsFromEvents(taxEvents []events.Event) []float64 {
	taxAmounts := make([]float64, len(taxEvents))

	for index, taxEvent := range taxEvents {
		taxAmounts[index] = taxEvent.Amount()
	}

	return taxAmounts
}

// ToJson converts any data structure to its JSON string representation.
func ToJson(data any) string {
	bytes, marshalError := json.Marshal(data)

	if marshalError != nil {
		panic(marshalError)
	}

	return string(bytes)
}
