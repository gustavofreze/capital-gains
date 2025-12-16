package outbound

import "capital-gains/src/application/domain/models"

// Operations represents the output boundary for storing market operations
// that will be used in a single capital gain calculation lifecycle.
type Operations interface {
	// Save persists a new market operation in the current calculation context.
	//
	// [param]  operation models.Operation      instance to be stored.
	Save(operation models.Operation)

	// FindAll returns all stored Operation aggregates and clears the storage,
	// so a subsequent call returns an empty list.
	//
	// [return] []models.Operation            list of stored aggregates.
	FindAll() []models.Operation
}
