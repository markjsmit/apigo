package definition

type Path struct {
	Tags        []string         `json:"tags,omitempty"`
	Summary     string           `json:"summary"`
	Description string           `json:"description,omitempty"`
	OperationID string           `json:"operationId,omitempty"`
	Consumes    []string         `json:"consumes,omitempty"`
	Produces    []string         `json:"produces,omitempty"`
	Parameters  []Parameter      `json:"parameters,omitempty"`
	Responses   map[int]Response `json:"responses,omitempty"`
	Deprecated  bool             `json:"deprecated"`
}

type Response struct {
	Description string `json:"description,omitempty"`
	Schema      Schema `json:"schema,omitempty"`
}

type Schema struct {
	Ref     string   `json:"$ref,omitempty"`
	Type    string   `json:"type,omitempty"`
	Items   []Schema `json:"items,omitempty"`
	Enum    []string `json:"enum,omitempty"`
	Default []string `json:"default,omitempty"`
}

type Parameter struct {
	Name             string `json:"name"`
	In               string `json:"in"`
	Description      string `json:"description"`
	Required         bool   `json:"required"`
	Type             string `json:"type"`
	Maximum          int    `json:"maximum,omitempty"`
	Minimum          int    `json:"minimum,omitempty"`
	CollectionFormat string `json:"collectionFormat,omitempty"`
	Format           string `json:"format,omitempty"`
	Schema           Schema `json:"schema,omitempty"`
}
