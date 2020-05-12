package definition


type Info struct {
	Description    string  `json:"description,omitempty"`
	Version        string  `json:"version,omitempty"`
	Title          string  `json:"title,omitempty"`
	TermsOfService string  `json:"termsOfService,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
	License        License `json:"license,omitempty"`
}

type Contact struct {
	Email string `json:"email,omitempty"`
}
type License struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}