package gosiren

type SirenEntity[P any] struct {
	Title      string            `json:"title,omitempty"`
	Class      []string          `json:"class,omitempty"`
	Properties *P                `json:"properties,omitempty"`
	Actions    []SirenAction     `json:"actions,omitempty"`
	Links      []SirenLink       `json:"links,omitempty"`
	Entities   []SirenEmbed[any] `json:"entities,omitempty"`
}

func (e *SirenEntity[P]) SetTitle(title string) *SirenEntity[P] {
	e.Title = title
	return e
}

func (e *SirenEntity[P]) SetClasses(classes []string) *SirenEntity[P] {
	e.Class = classes
	return e
}

func (e *SirenEntity[P]) AddClass(class string) *SirenEntity[P] {
	if e.Class == nil {
		e.Class = []string{}
	}
	e.Class = append(e.Class, class)
	return e
}

func (e *SirenEntity[P]) SetProperties(props *P) *SirenEntity[P] {
	e.Properties = props
	return e
}

func (e *SirenEntity[P]) SetLinks(links []SirenLink) *SirenEntity[P] {
	e.Links = links
	return e
}

func (e *SirenEntity[P]) AddLink(rel []string, href string, l SirenLink) *SirenEntity[P] {
	l.Rel = rel
	l.Href = href
	e.Links = append(e.Links, l)
	return e
}

func (e *SirenEntity[P]) SetActions(actions []SirenAction) *SirenEntity[P] {
	e.Actions = actions
	return e
}

func (e *SirenEntity[P]) AddAction(name string, title string, method string,
	href string, ctype string, fields ...SirenField) *SirenEntity[P] {
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
func (e *SirenEntity[P]) AddEmbed(href string, siren *SirenEntity[any], rel ...string) *SirenEntity[P] {
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

func NewSirenEntity[P any]() *SirenEntity[P] {
	return &SirenEntity[P]{
		Class:    []string{},
		Entities: []SirenEmbed[any]{},
		Actions:  []SirenAction{},
		Links:    []SirenLink{},
	}
}
