package events

// Event defines the domain representation of a tax-related outcome produced
// when processing a financial operation in a capital gain calculation.
type Event interface {
	// Amount returns the tax amount associated with this event.
	// A zero amount represents that no tax is due for the operation.
	//
	// [return] float64   tax amount associated with this event.
	Amount() float64
}
