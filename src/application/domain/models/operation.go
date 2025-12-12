package models

// Operation defines the domain representation of a financial market transaction
// that affects the investor position in the capital gains' context.
type Operation interface {
	// Quantity returns how many units (for example, shares) are traded in the operation.
	//
	// [return] Quantity   number of units traded in the operation.
	Quantity() Quantity

	// TotalValue calculates the total monetary value of the operation,
	// which is the product of the unit cost and the quantity.
	//
	// [return] MonetaryValue   total value of the operation.
	TotalValue() MonetaryValue
}
