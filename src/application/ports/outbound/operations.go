package outbound

import "capital-gains/src/application/domain/models"

// Operations represents the output boundary for storing market operations
// that will be used in a single capital gain calculation lifecycle.
type Operations interface {
	// Save persists a new market operation in the current calculation context.
	//
	// [param]  operation models.Operation      instance to be stored.
	// [return] error                           non-nil error if the operation cannot be stored.
	Save(operation models.Operation) error

	// FindAll returns all stored market operations in the order they were added,
	// to be processed by the capital gain aggregate.
	//
	// [return] []models.Operation              list of operations registered for the current calculation.
	FindAll() []models.Operation
}
