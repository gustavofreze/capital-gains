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
	Save(capitalGain models.CapitalGain)

	// FindAll returns all stored CapitalGain aggregates and clears the storage,
	// so a subsequent call returns an empty list.
	//
	// [return] []models.CapitalGain            list of stored aggregates.
	FindAll() []models.CapitalGain
}
