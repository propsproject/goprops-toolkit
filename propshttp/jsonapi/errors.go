package jsonapi

type Error struct {
	Code   string      `json:"code"`
	Detail string      `json:"detail"`
	Source interface{} `json:"source"`
}

func (e *Error) SetCode(code string) *Error {
	e.Code = code
	return e
}

func (e *Error) SetDetail(detail string) *Error {
	e.Detail = detail
	return e
}

func (e *Error) SetSource(source interface{}) *Error {
	e.Source = source
	return e
}

type Errors []*Error

func ResponseError() *Error {
	return &Error{}
}

