package driver

type Request struct {
	operations []Operation
}

func NewRequest(operations []Operation) Request {
	return Request{operations: operations}
}

func (request *Request) Operations() []Operation {
	return request.operations
}
