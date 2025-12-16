package capitalgains

import "capital-gains/src/application/domain/models"

type Repository struct {
	capitalGains []models.CapitalGain
}

func NewRepository() *Repository {
	return &Repository{
		capitalGains: make([]models.CapitalGain, 0),
	}
}

func (repository *Repository) Save(capitalGain models.CapitalGain) {
	repository.capitalGains = append(repository.capitalGains, capitalGain)
}

func (repository *Repository) FindAll() []models.CapitalGain {
	capitalGains := make([]models.CapitalGain, len(repository.capitalGains))
	copy(capitalGains, repository.capitalGains)

	repository.capitalGains = make([]models.CapitalGain, 0)

	return capitalGains
}
