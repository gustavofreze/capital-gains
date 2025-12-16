package test

import "capital-gains/src/application/commands"

type RegisterBuyHandlerMock struct {
	calls    int
	received []commands.RegisterBuy
}

func NewRegisterBuyHandlerMock() *RegisterBuyHandlerMock {
	return &RegisterBuyHandlerMock{
		calls:    0,
		received: []commands.RegisterBuy{},
	}
}

func (mock *RegisterBuyHandlerMock) Handle(command commands.RegisterBuy) {
	mock.calls++
	mock.received = append(mock.received, command)
}

func (mock *RegisterBuyHandlerMock) Calls() int {
	return mock.calls
}

func (mock *RegisterBuyHandlerMock) FirstReceived() commands.RegisterBuy {
	return mock.received[0]
}
