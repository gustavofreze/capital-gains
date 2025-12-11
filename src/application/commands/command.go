package commands

// Command defines the inbound representation of a financial market transaction
// that will be transformed into a domain operation in the capital gain context.
type Command interface {
	// UnitCost returns the price per unit (for example, per share) used in the command.
	//
	// [return] float64   unit cost specified by the command.
	UnitCost() float64

	// Quantity returns how many units (for example, shares) are traded in the command.
	//
	// [return] int64   number of units specified by the command.
	Quantity() int64
}
