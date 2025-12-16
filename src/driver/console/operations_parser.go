package console

import (
	"capital-gains/src/driver"
	"encoding/json"
	"strings"
)

type OperationsParser struct{}

func NewOperationsParser() *OperationsParser {
	return new(OperationsParser)
}

func (parser *OperationsParser) Parse(payload string) (driver.Request, bool) {
	trimmedPayload := strings.TrimSpace(payload)

	if trimmedPayload == "" {
		return driver.Request{}, false
	}

	var operations []driver.Operation

	err := json.Unmarshal([]byte(trimmedPayload), &operations)

	if err != nil {
		return driver.Request{}, false
	}

	return driver.NewRequest(operations), true
}
