package test

import "capital-gains/src/application/commands"

type RegisterSellHandlerMock struct {
	calls    int
	received []commands.RegisterSell
}

func NewRegisterSellHandlerMock() *RegisterSellHandlerMock {
	return &RegisterSellHandlerMock{
		calls:    0,
		received: []commands.RegisterSell{},
	}
}

func (mock *RegisterSellHandlerMock) Handle(command commands.RegisterSell) {
	mock.calls++
	mock.received = append(mock.received, command)
}

func (mock *RegisterSellHandlerMock) Calls() int {
	return mock.calls
}

func (mock *RegisterSellHandlerMock) FirstReceived() commands.RegisterSell {
	return mock.received[0]
}
