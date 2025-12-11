package inbound

import "capital-gains/src/application/commands"

// CommandHandler defines the input boundary responsible for receiving
// financial transaction commands and delegating them to the application layer.
type CommandHandler interface {
	// Handle processes a single financial transaction command within the
	// current capital gain calculation lifecycle.
	//
	// [param]  command commands.Command   instance to be handled.
	// [return] error                      non-nil error if the command cannot be processed.
	Handle(command commands.Command) error
}
