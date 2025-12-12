package operations

import (
	"capital-gains/src/application/domain/models"
	"capital-gains/src/application/ports/outbound"
)

var _ outbound.Operations = (*Repository)(nil)

type Repository struct {
	operations []models.Operation
}

func NewRepository() *Repository {
	return &Repository{
		operations: make([]models.Operation, 0),
	}
}

func (repository *Repository) Save(operation models.Operation) {
	repository.operations = append(repository.operations, operation)
}

func (repository *Repository) FindAll() []models.Operation {
	operations := make([]models.Operation, len(repository.operations))
	copy(operations, repository.operations)

	return operations
}
