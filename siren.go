package gosiren

type SirenLink struct {
	Rel  []string `json:"rel,omitempty"`
	Href string   `json:"href,omitempty"`
}

type SirenAction struct {
	Name   string       `json:"name"`
	Title  string       `json:"title"`
	Method string       `json:"method"`
	Href   string       `json:"href"`
	Type   string       `json:"type"`
	Fields []SirenField `json:"fields"`
}

type SirenField struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type SirenEntity struct {
	Class      []string               `json:"class"`
	Properties map[string]interface{} `json:"properties"`
	Entities   []SirenEntity          `json:"entities"`
	Actions    []SirenAction          `json:"actions"`
	Links      []SirenLink            `json:"links"`
	Title      string                 `json:"title"`

	// Optional for root, required for embedded
	SirenLink
}

func NewSirenEntity() *SirenEntity {
	return &SirenEntity{
		Class:      []string{},
		Properties: map[string]interface{}{},
		Entities:   []SirenEntity{},
		Actions:    []SirenAction{},
		Links:      []SirenLink{},
	}
}
