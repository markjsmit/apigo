package config

type DocsConfig struct {
	Title                   string
	Email                   string
	Description             string
	LicenseName             string
	LicenseUrl              string
	TermsOfService          string
	Version                 string
	BasePath                string
	ExternalDocsDescription string
	ExternalDocsUrl         string
	Host                    string
	Schemes                 []string
}

func NewDocsConfig() *DocsConfig {
	return &DocsConfig{
		Title:                   "Example docs",
		Version:                 "1.0",
		TermsOfService:          "",
		Description:             "An automatic generated api by apigo.",
		Email:                   "markjsmit@hotmail.com",
		LicenseName:             "MIT",
		LicenseUrl:              "",
		BasePath:                "/",
		Host:                    "localhost:8088",
		Schemes:                 []string{"http"},
		ExternalDocsDescription: "External documentation",
		ExternalDocsUrl:         "http://test.docs.com/",
	};
}
