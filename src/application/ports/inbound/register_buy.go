package inbound

import "capital-gains/src/application/commands"

// RegisterBuy defines the input boundary responsible for handling buy
// operations and updating the portfolio state for the current lifecycle.
type RegisterBuy interface {
	// Handle registers a buy operation, updating the investor position
	// based on the provided command.
	//
	// [param]  command commands.RegisterBuy   buy operation command to be handled.
	Handle(command commands.RegisterBuy)
}
