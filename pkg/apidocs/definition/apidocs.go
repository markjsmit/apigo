package definition

type ApiDocs struct {
	Swagger      string                     `json:"swagger"`
	Info         Info                       `json:"info"`
	Host         string                     `json:"host"`
	BasePath     string                     `json:"basePath"`
	Tags         []Tags                     `json:"tags"`
	Schemes      []string                   `json:"schemes"`
	Paths        map[string]map[string]Path `json:"paths"`
	Definitions  map[string]Definition      `json:"definitions"`
	ExternalDocs ExternalDocs               `json:"externalDocs"`
}

func NewApiDocs() *ApiDocs {
	return &ApiDocs{}
}
