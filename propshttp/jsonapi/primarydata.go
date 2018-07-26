package jsonapi

//The document’s “primary data” is a representation of the resource or collection of resources targeted by a request.
type PrimaryData struct {
	//Primary data MUST be either:
	//a single resource identifier object
	Resource map[string]interface{} `json:"resource,omitempty"`
}

func (p *PrimaryData) SetResource(key string, resource *Resource) *PrimaryData {
	p.Resource[key] = resource
	return p
}