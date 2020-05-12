package config

type EntityConfig struct {
	PageSize                int
	Description             string
	Name                    string
	Path                    string
	ExternalDocsUrl         string
	ExternalDocsDescription string
	IsDeprecated            bool
	Consumes                []string
	Produces                []string
	PageParam               string
}

func NewEntityConfig() *EntityConfig {
	return &EntityConfig{
		PageSize:                30,
		PageParam:               "page",
		Path:                    "/api/__name__",
		Name:                    "__name__",
		ExternalDocsUrl:         "__docs__/__name__",
		ExternalDocsDescription: "Find more about __name__",
		Description:             "Everything related to __name__",
		IsDeprecated:            false,
		Consumes:                []string{"application/json"},
		Produces:                []string{"application/json"},
	};
}
