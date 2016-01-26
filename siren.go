package gosiren

type SirenLink struct {
	Rel  []string `json:"rel,omitempty"`
	Href string   `json:"href,omitempty"`
}

type SirenAction struct {
	Name   string       `json:"name"`
	Href   string       `json:"href"`
	Title  string       `json:"title,omitempty"`
	Method string       `json:"method,omitempty"`
	Type   string       `json:"type,omitempty"`
	Fields []SirenField `json:"fields,omitempty"`
}

type SirenField struct {
	Name  string `json:"name"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type"`
	Value string `json:"value,omitempty"`
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

func (e *SirenEntity) AddLink(href string, rel ...string) {
	e.Links = append(e.Links, SirenLink{
		Rel:  rel,
		Href: href,
	})
}

func (e *SirenEntity) AddAction(name string, title string, method string,
	href string, ctype string, fields ...SirenField) {
	e.Actions = append(e.Actions, SirenAction{
		Name:   name,
		Title:  title,
		Method: method,
		Href:   href,
		Type:   ctype,
		Fields: fields,
	})
}

func MakeField(name string, vtype string, value string) SirenField {
	return SirenField{
		Name:  name,
		Type:  vtype,
		Value: value,
	}
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
