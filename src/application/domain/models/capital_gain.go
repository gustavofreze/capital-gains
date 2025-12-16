package models

import "capital-gains/src/application/domain/events"

type CapitalGain struct {
	events   []events.Event
	position Position
}

func NewCapitalGain() CapitalGain {
	return CapitalGain{
		events:   make([]events.Event, 0),
		position: NewPosition(),
	}
}

func (capitalGain *CapitalGain) ApplyOperations(operations []Operation) {
	for _, operation := range operations {
		tax := operation.ApplyTo(&capitalGain.position)

		if tax.IsExempted() {
			capitalGain.events = append(capitalGain.events, events.NewTaxExempted())
			continue
		}

		capitalGain.events = append(
			capitalGain.events,
			events.NewTaxPaid(tax.Value().ToFloat64()),
		)
	}
}

func (capitalGain *CapitalGain) Events() []events.Event {
	taxEvents := make([]events.Event, len(capitalGain.events))
	copy(taxEvents, capitalGain.events)

	return taxEvents
}
