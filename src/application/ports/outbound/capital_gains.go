package outbound

import "capital-gains/src/application/domain/models"

// CapitalGains defines the output boundary responsible for storing and
// retrieving CapitalGain aggregates, which encapsulate the investor
// position and all tax events produced in a calculation lifecycle.
type CapitalGains interface {
	// Save persists the final state of a CapitalGain aggregate for the
	// current calculation lifecycle.
	//
	// [param]  capitalGain models.CapitalGain   aggregate instance to be stored.
	// [return] error                            non-nil error if the aggregate cannot be stored.
	Save(capitalGain models.CapitalGain) error

	// FindAll returns all stored CapitalGain aggregates.
	//
	// [return] []models.CapitalGain            list of stored aggregates.
	FindAll() []models.CapitalGain
}
