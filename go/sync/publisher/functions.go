package publisher

// return false if the provided address is to be filtered out
type AddressFilter func(address string) bool

type SendCallback func(address string, response *Response)
