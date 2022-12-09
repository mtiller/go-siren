package siren

type EmbeddedEntity[P any] struct {
	Rel []string `json:"rel"`

	// Only present if this is an "embedded link" (href required in that case)
	Href string `json:"href,omitempty"`
	Type string `json:"type,omitempty"`

	// Otherwise, just like a normal SirenEntity
	Title      string                `json:"title,omitempty"`
	Class      []string              `json:"class,omitempty"`
	Properties P                     `json:"properties,omitempty"`
	Entities   []EmbeddedEntity[any] `json:"entities,omitempty"`
	Actions    []SirenAction         `json:"actions,omitempty"`
	Links      []SirenLink           `json:"links,omitempty"`
}

func (e *EmbeddedEntity[P]) GetHref() string {
	return e.Href
}

func (e *EmbeddedEntity[P]) SetHref(href string) *EmbeddedEntity[P] {
	e.Href = href
	return e
}

func (e *EmbeddedEntity[P]) GetRel() []string {
	return e.Rel
}

func (e *EmbeddedEntity[P]) AddRel(rel string) *EmbeddedEntity[P] {
	if e.Rel == nil {
		e.Rel = []string{}
	}
	e.Rel = append(e.Rel, rel)
	return e
}

func (e *EmbeddedEntity[P]) SetRel(rels []string) *EmbeddedEntity[P] {
	e.Rel = rels
	return e
}

func (e *EmbeddedEntity[P]) SetType(t string) *EmbeddedEntity[P] {
	e.Type = t
	return e
}

func (e *EmbeddedEntity[P]) SetTitle(title string) *EmbeddedEntity[P] {
	e.Title = title
	return e
}

func (e *EmbeddedEntity[P]) SetClasses(classes []string) *EmbeddedEntity[P] {
	e.Class = classes
	return e
}

func (e *EmbeddedEntity[P]) AddClass(class string) *EmbeddedEntity[P] {
	if e.Class == nil {
		e.Class = []string{}
	}
	e.Class = append(e.Class, class)
	return e
}

func (e *EmbeddedEntity[P]) SetProperties(props P) *EmbeddedEntity[P] {
	e.Properties = props
	return e
}

func (e *EmbeddedEntity[P]) SetLinks(links []SirenLink) *EmbeddedEntity[P] {
	e.Links = links
	return e
}

func (e *EmbeddedEntity[P]) AddLink(rel []string, href string, l SirenLink) *EmbeddedEntity[P] {
	l.Rel = rel
	l.Href = href
	e.Links = append(e.Links, l)
	return e
}

func (e *EmbeddedEntity[P]) SetActions(actions []SirenAction) *EmbeddedEntity[P] {
	e.Actions = actions
	return e
}

func (e *EmbeddedEntity[P]) AddAction(name string, title string, method string,
	href string, ctype string, fields ...SirenField) *EmbeddedEntity[P] {
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

func (e *EmbeddedEntity[P]) AddEntity(siren *Siren[any], rel ...string) *EmbeddedEntity[P] {
	embed := EmbeddedEntity[any]{
		Rel:        rel,
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

func NewSirenEmbed[P any](rel []string) *EmbeddedEntity[P] {
	return &EmbeddedEntity[P]{
		Rel:      rel,
		Class:    []string{},
		Entities: []EmbeddedEntity[any]{},
		Actions:  []SirenAction{},
		Links:    []SirenLink{},
	}
}
