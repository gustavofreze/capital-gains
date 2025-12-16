package test

import "capital-gains/src/application/commands"

type CalculateCapitalGainHandlerMock struct {
	calls    int
	received []commands.CalculateCapitalGain
}

func NewCalculateCapitalGainHandlerMock() *CalculateCapitalGainHandlerMock {
	return &CalculateCapitalGainHandlerMock{
		calls:    0,
		received: []commands.CalculateCapitalGain{},
	}
}

func (mock *CalculateCapitalGainHandlerMock) Handle(command commands.CalculateCapitalGain) {
	mock.calls++
	mock.received = append(mock.received, command)
}

func (mock *CalculateCapitalGainHandlerMock) Calls() int {
	return mock.calls
}

func (mock *CalculateCapitalGainHandlerMock) FirstReceived() commands.CalculateCapitalGain {
	return mock.received[0]
}
