package test

import "capital-gains/src/driver/console"

var _ console.Console = (*ConsoleMock)(nil)

type ConsoleMock struct {
	inputs  []string
	outputs []string
}

func NewConsoleMock(inputs []string) *ConsoleMock {
	return &ConsoleMock{
		inputs:  inputs,
		outputs: make([]string, 0),
	}
}

func (console *ConsoleMock) ReadLines() []string {
	return console.inputs
}

func (console *ConsoleMock) WriteLine(line string) {
	console.outputs = append(console.outputs, line)
}

func (console *ConsoleMock) WrittenLines() []string {
	return console.outputs
}

func (console *ConsoleMock) GetByIndex(index int) string {
	if len(console.outputs) == index {
		return ""
	}

	return console.outputs[index]
}
