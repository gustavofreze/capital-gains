package commandbus_test

import (
	"testing"

	"capital-gains/src/driver/commandbus"

	"capital-gains/test"

	"capital-gains/src/application/commands"
	"capital-gains/src/driver/console"

	"github.com/stretchr/testify/assert"
)

func TestCommandMapperMapGivenRequestWithOperationsWhenMapThenReturnsCommandsInSameOrder(t *testing.T) {
	t.Parallel()

	// Given a JSON payload representing a list of market operations
	payload := []map[string]any{
		{"operation": "buy", "unit-cost": 10.00, "quantity": 100},
		{"operation": "sell", "unit-cost": 15.00, "quantity": 50},
		{"operation": "buy", "unit-cost": 20.00, "quantity": 10},
	}
	payloadJSON := test.ToJson(payload)

	// And I parse the payload into a request
	parser := console.NewOperationsParser()
	request, ok := parser.Parse(payloadJSON)
	assert.True(t, ok)

	// And I create a command mapper for this request
	mapper := commandbus.NewCommandMapper(request)

	// When I map the request operations into commands
	mappedCommands := mapper.Map()

	// Then I expect one command per operation
	assert.Len(t, mappedCommands, 3)

	// And I expect the mapped commands to match each operation command conversion and preserve the same order
	operations := request.Operations()
	assert.Equal(t, operations[0].ToCommand(), mappedCommands[0])
	assert.Equal(t, operations[1].ToCommand(), mappedCommands[1])
	assert.Equal(t, operations[2].ToCommand(), mappedCommands[2])

	// And I expect the mapped command types to match the original operations (buy, sell, buy)
	_, isFirstBuy := mappedCommands[0].(commands.RegisterBuy)
	assert.True(t, isFirstBuy)

	_, isSell := mappedCommands[1].(commands.RegisterSell)
	assert.True(t, isSell)

	_, isSecondBuy := mappedCommands[2].(commands.RegisterBuy)
	assert.True(t, isSecondBuy)
}

func TestCommandMapperMapGivenRequestWithNoOperationsWhenMapThenReturnsEmptyList(t *testing.T) {
	t.Parallel()

	// Given a JSON payload representing an empty list of operations
	var payload []map[string]any
	payloadJSON := test.ToJson(payload)

	// And I parse the payload into a request
	parser := console.NewOperationsParser()
	request, ok := parser.Parse(payloadJSON)
	assert.True(t, ok)

	// And I create a command mapper for this request
	mapper := commandbus.NewCommandMapper(request)

	// When I map the request operations into commands
	mappedCommands := mapper.Map()

	// Then I expect an empty command list
	assert.NotNil(t, mappedCommands)
	assert.Len(t, mappedCommands, 0)
}
