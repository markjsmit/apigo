package definition

type Definition struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
}

type Property struct {
	Ref        string              `json:"$ref,omitempty"`
	Type       string              `json:"type,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
	Example    string              `json:"example,omitempty"`
}
