package gosiren

type SirenEmbed[P any] struct {
	Rel []string `json:"rel"`

	// Only present if this is an "embedded link" (href required in that case)
	Href string `json:"href,omitempty"`
	Type string `json:"type,omitempty"`

	// Otherwise, just like a normal SirenEntity
	Title      string            `json:"title,omitempty"`
	Class      []string          `json:"class,omitempty"`
	Properties *P                `json:"properties,omitempty"`
	Entities   []SirenEmbed[any] `json:"entities,omitempty"`
	Actions    []SirenAction     `json:"actions,omitempty"`
	Links      []SirenLink       `json:"links,omitempty"`
}

func (e *SirenEmbed[P]) GetHref() string {
	return e.Href
}

func (e *SirenEmbed[P]) SetHref(href string) *SirenEmbed[P] {
	e.Href = href
	return e
}

func (e *SirenEmbed[P]) GetRel() []string {
	return e.Rel
}

func (e *SirenEmbed[P]) AddRel(rel string) *SirenEmbed[P] {
	if e.Rel == nil {
		e.Rel = []string{}
	}
	e.Rel = append(e.Rel, rel)
	return e
}

func (e *SirenEmbed[P]) SetRel(rels []string) *SirenEmbed[P] {
	e.Rel = rels
	return e
}

func (e *SirenEmbed[P]) SetType(t string) *SirenEmbed[P] {
	e.Type = t
	return e
}

func (e *SirenEmbed[P]) SetTitle(title string) *SirenEmbed[P] {
	e.Title = title
	return e
}

func (e *SirenEmbed[P]) SetClasses(classes []string) *SirenEmbed[P] {
	e.Class = classes
	return e
}

func (e *SirenEmbed[P]) AddClass(class string) *SirenEmbed[P] {
	if e.Class == nil {
		e.Class = []string{}
	}
	e.Class = append(e.Class, class)
	return e
}

func (e *SirenEmbed[P]) SetProperties(props *P) *SirenEmbed[P] {
	e.Properties = props
	return e
}

func (e *SirenEmbed[P]) SetLinks(links []SirenLink) *SirenEmbed[P] {
	e.Links = links
	return e
}

func (e *SirenEmbed[P]) AddLink(rel []string, href string, l SirenLink) *SirenEmbed[P] {
	l.Rel = rel
	l.Href = href
	e.Links = append(e.Links, l)
	return e
}

func (e *SirenEmbed[P]) SetActions(actions []SirenAction) *SirenEmbed[P] {
	e.Actions = actions
	return e
}

func (e *SirenEmbed[P]) AddAction(name string, title string, method string,
	href string, ctype string, fields ...SirenField) *SirenEmbed[P] {
	e.Actions = append(e.Actions, SirenAction{
		Name:   name,
		Title:  title,
		Method: method,
		Href:   href,
		Type:   ctype,
		Fields: fields,
	})
	return e
}

// Rel is an argument because it is required
func (e *SirenEmbed[P]) AddEmbed(href string, siren *SirenEntity[any], rel ...string) *SirenEmbed[P] {
	embed := SirenEmbed[any]{
		Rel:        rel,
		Href:       href,
		Title:      siren.Title,
		Class:      siren.Class,
		Properties: siren.Properties,
		Entities:   siren.Entities,
		Actions:    siren.Actions,
		Links:      siren.Links,
	}
	e.Entities = append(e.Entities, embed)
	return e
}

func NewSirenEmbed[P any](rel []string) *SirenEmbed[P] {
	return &SirenEmbed[P]{
		Rel:      rel,
		Class:    []string{},
		Entities: []SirenEmbed[any]{},
		Actions:  []SirenAction{},
		Links:    []SirenLink{},
	}
}
