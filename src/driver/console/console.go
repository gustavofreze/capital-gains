package console

// Console abstracts the terminal I/O used by the application.
type Console interface {
	// ReadLines returns all lines read from the standard input until EOF.
	ReadLines() []string

	// WriteLine writes a single line to the standard output.
	WriteLine(text string)
}
