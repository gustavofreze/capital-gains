package console

import (
	"capital-gains/src/driver"
	"strings"
)

type OperationsConsole struct {
	console Console
	parser  *OperationsParser
}

func NewOperationsConsole(console Console) *OperationsConsole {
	return &OperationsConsole{
		console: console,
		parser:  NewOperationsParser(),
	}
}

func (operationsConsole *OperationsConsole) ReadRequests() []driver.Request {
	lines := operationsConsole.console.ReadLines()
	completePayload := strings.TrimSpace(strings.Join(lines, "\n"))

	if request, ok := operationsConsole.parser.Parse(completePayload); ok {
		return []driver.Request{request}
	}

	requests := make([]driver.Request, 0)

	for _, line := range lines {
		linePayload := strings.TrimSpace(line)

		if linePayload == "" {
			continue
		}

		request, ok := operationsConsole.parser.Parse(linePayload)

		if !ok {
			panic("invalid input: expected a JSON array of operations")
		}

		requests = append(requests, request)
	}

	return requests
}

func (operationsConsole *OperationsConsole) WriteResponse(response driver.Response) {
	operationsConsole.console.WriteLine(response.ToString())
}
