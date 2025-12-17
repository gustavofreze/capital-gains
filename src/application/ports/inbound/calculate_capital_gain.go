package inbound

import "capital-gains/src/application/commands"

// CalculateCapitalGain defines the input boundary responsible for executing
// the capital gains calculation use case based on a CalculateCapitalGain command.
type CalculateCapitalGain interface {
	// Handle executes the capital gains calculation for the current input
	// lifecycle, producing the tax outcomes derived from the provided operations.
	//
	// [param]  command commands.CalculateCapitalGain   use case command to be handled.
	Handle(command commands.CalculateCapitalGain)
}
