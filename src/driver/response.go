package driver

import (
	"encoding/json"

	"capital-gains/src/application/domain/models"
)

type Response struct {
	taxes []Tax
}

func NewResponse(capitalGains []models.CapitalGain) Response {
	taxItems := make([]Tax, 0)

	for _, capitalGain := range capitalGains {
		taxEvents := capitalGain.Events()

		for _, taxEvent := range taxEvents {
			taxItems = append(taxItems, NewTax(taxEvent.Amount()))
		}
	}

	return Response{taxes: taxItems}
}

func (response Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(response.taxes)
}

func (response Response) ToString() string {
	serializedResponse, err := json.Marshal(response)

	if err != nil {
		panic(err)
	}

	return string(serializedResponse)
}
