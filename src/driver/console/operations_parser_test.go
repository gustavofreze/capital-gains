package console_test

import (
	"capital-gains/src/driver/console"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperationsParserParsesValidJsonArray(t *testing.T) {
	t.Parallel()

	// Given a valid JSON array representing market operations
	payload := `[{"operation":"buy","unit-cost":10.00,"quantity":100}]`
	parser := console.NewOperationsParser()

	// When parsing the payload
	request, ok := parser.Parse(payload)

	// Then I expect the parsing to succeed
	assert.True(t, ok)

	// And I expect a request with a single operation
	assert.Len(t, request.Operations(), 1)
}

func TestOperationsParserReturnsFalseWhenPayloadIsNotValidJsonArray(t *testing.T) {
	t.Parallel()

	// Given an invalid payload that is not a JSON array of operations
	payload := `this is not json`
	parser := console.NewOperationsParser()

	// When parsing the payload
	_, ok := parser.Parse(payload)

	// Then I expect the parsing to fail
	assert.False(t, ok)
}

func TestOperationsParserReturnsFalseWhenPayloadIsEmpty(t *testing.T) {
	t.Parallel()

	// Given an empty payload
	payload := `   `
	parser := console.NewOperationsParser()

	// When parsing the payload
	_, ok := parser.Parse(payload)

	// Then I expect the parsing to fail
	assert.False(t, ok)
}
