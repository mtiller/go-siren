package gosiren

type SirenLink struct {
	Rel   []string `json:"rel"`
	Href  string   `json:"href"`
	Class []string `json:"class,omitempty"`
	Type  string   `json:"type,omitempty"`
	Title string   `json:"title,omitempty"`
}
