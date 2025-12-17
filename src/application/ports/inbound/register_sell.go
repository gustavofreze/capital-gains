package inbound

import "capital-gains/src/application/commands"

// RegisterSell defines the input boundary responsible for handling sell
// operations and producing the corresponding tax event for the current lifecycle.
type RegisterSell interface {
	// Handle registers a sell operation, updating the investor position and
	// computing the resulting tax event based on the provided command.
	//
	// [param]  command commands.RegisterSell   sell operation command to be handled.
	Handle(command commands.RegisterSell)
}
