package models

// Operation defines the domain representation of a financial market transaction
// that affects the investor position in the capital gains' context.
type Operation interface {
	// ApplyTo applies this operation to the given position and returns the resulting tax.
	//
	// [param]  position *Position   current investor position to be updated.
	//
	// [return] Tax   produced by this operation (zero means exempted).
	ApplyTo(position *Position) Tax
}
