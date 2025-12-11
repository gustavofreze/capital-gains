package models

// Operation defines the domain representation of a financial market transaction
// that affects the investor position in the capital gains' context.
type Operation interface {
	// UnitCost returns the price per unit (for example, per share) used in the operation.
	//
	// [return] UnitCost   unit cost used in the operation.
	UnitCost() UnitCost

	// Quantity returns how many units (for example, shares) are traded in the operation.
	//
	// [return] Quantity   number of units traded in the operation.
	Quantity() Quantity
}
