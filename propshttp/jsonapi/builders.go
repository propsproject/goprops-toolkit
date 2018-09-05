package jsonapi

type AllBuilders struct {
	Response    ResBuilder
}

func Builders() *AllBuilders {
	return &AllBuilders{
		Response:    ResponseBuilder(),
	}
}