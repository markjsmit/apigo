package definition

type Tags struct {
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	ExternalDocs ExternalDocs `json:"externalDocs,omitempty"`
}
