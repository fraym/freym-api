package publisher

type Response struct {
	ack                      bool
	requestedNodeUnavailable bool
}

func NewResponse(ack bool, requestedNodeUnavailable bool) *Response {
	return &Response{
		ack:                      ack,
		requestedNodeUnavailable: requestedNodeUnavailable,
	}
}

func (r *Response) IsAcknowledged() bool {
	return r.ack
}

func (r *Response) IsRequestedNodeUnavailable() bool {
	return r.requestedNodeUnavailable
}
