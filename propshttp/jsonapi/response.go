package jsonapi

type Response struct {
	//A document MUST contain at least one of the following top-level members
	Data        *PrimaryData `json:"data,omitempty"`   //“primary data” The members data and errors MUST NOT coexist in the same document.
	Errors      Errors      `json:"errors,omitempty"` //an array of error objects
	Meta        *Meta        `json:"meta,omitempty"`   //a meta object that contains non-standard meta-information.
	ContentType string       `json:"-"`
	Status      int          `json:"-"`
}
