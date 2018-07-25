package rediscache

type Encoder interface {
	Encode() ([]byte, error)
}

func NewEncoder(e interface{}) Encoder {
	return e.(Encoder)
}

type Decoder interface {
	Decode([]byte, interface{}) error
}

func NewDecoder(d interface{}) Decoder {
	return d.(Decoder)
}