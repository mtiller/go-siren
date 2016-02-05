package gosiren

type Siren interface {
	SetTitle(string)
	SetClasses([]string)
	SetProperties(map[string]interface{})
	SetEntities([]SirenEmbed)
	SetActions([]SirenAction)
	SetLinks([]SirenLink)
}

type SirenEntity struct {
	Title      string                 `json:"title,omitempty"`
	Class      []string               `json:"class,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Entities   []SirenEmbed           `json:"entities,omitempty"`
	Actions    []SirenAction          `json:"actions,omitempty"`
	Links      []SirenLink            `json:"links,omitempty"`
}

type SirenLink struct {
	Rel   []string `json:"rel"`
	Href  string   `json:"href"`
	Class []string `json:"class,omitempty"`
	Type  string   `json:"type,omitempty"`
	Title string   `json:"title,omitempty"`
}

type SirenAction struct {
	Name   string       `json:"name"`
	Href   string       `json:"href"`
	Class  []string     `json:"class,omitempty"`
	Method string       `json:"method,omitempty"`
	Title  string       `json:"title,omitempty"`
	Type   string       `json:"type,omitempty"`
	Fields []SirenField `json:"fields,omitempty"`
}

type SirenField struct {
	Name  string      `json:"name"`
	Class []string    `json:"class,omitempty"`
	Title string      `json:"title,omitempty"`
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type SirenEmbed struct {
	Rel []string `json:"rel"`

	// Only present if this is an "embedded link" (href required in that case)
	Href string `json:"href,omitempty"`
	Type string `json:"type,omitempty"`

	// Otherwise, just like a normal SirenEntity
	Title      string                 `json:"title,omitempty"`
	Class      []string               `json:"class,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Entities   []SirenEmbed           `json:"entities,omitempty"`
	Actions    []SirenAction          `json:"actions,omitempty"`
	Links      []SirenLink            `json:"links,omitempty"`
}

// Rel is an argument because it is required
func (e *SirenEntity) AddEmbed(rel []string, embed SirenEmbed) {
	embed.Rel = rel
	e.Entities = append(e.Entities, embed)
}

func (e *SirenEntity) AddLink(rel []string, href string, l SirenLink) {
	l.Rel = rel
	l.Href = href
	e.Links = append(e.Links, l)
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
		Entities:   []SirenEmbed{},
		Actions:    []SirenAction{},
		Links:      []SirenLink{},
	}
}
