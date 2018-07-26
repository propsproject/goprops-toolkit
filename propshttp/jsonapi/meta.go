package jsonapi

type Meta struct {
	//The value of each meta member MUST be an object (a “meta object”).
	Errors Errors `json:"errors"`
}

