package models

import "capital-gains/src/application/domain/events"

type CapitalGain struct {
	events   []events.Event
	position Position
}

func NewCapitalGains() CapitalGain {
	return CapitalGain{
		events:   make([]events.Event, 0),
		position: NewPosition(),
	}
}

func (capitalGain *CapitalGain) ApplyOperations(operations []Operation) {
	for _, capitalGainOperation := range operations {
		buyOperation, isBuyOperation := capitalGainOperation.(Buy)

		if isBuyOperation {
			capitalGain.position.ApplyBuy(buyOperation)
			taxEvent := events.NewTaxExempted()
			capitalGain.events = append(capitalGain.events, taxEvent)
			continue
		}

		sellOperation, isSellOperation := capitalGainOperation.(Sell)

		if isSellOperation {
			taxAmount := capitalGain.position.ApplySell(sellOperation)

			if taxAmount.IsZero() {
				taxEvent := events.NewTaxExempted()
				capitalGain.events = append(capitalGain.events, taxEvent)
				continue
			}

			taxEvent := events.NewTaxPaid(taxAmount.ToFloat64())
			capitalGain.events = append(capitalGain.events, taxEvent)
		}
	}
}

func (capitalGain *CapitalGain) Events() []events.Event {
	taxEvents := make([]events.Event, len(capitalGain.events))
	copy(taxEvents, capitalGain.events)

	return taxEvents
}
