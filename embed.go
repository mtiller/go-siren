package gosiren

type SirenEmbed[P any] struct {
	Rel []string `json:"rel"`

	// Only present if this is an "embedded link" (href required in that case)
	Href string `json:"href,omitempty"`
	Type string `json:"type,omitempty"`

	// Otherwise, just like a normal SirenEntity
	Title      string            `json:"title,omitempty"`
	Class      []string          `json:"class,omitempty"`
	Properties P                 `json:"properties,omitempty"`
	Entities   []SirenEmbed[any] `json:"entities,omitempty"`
	Actions    []SirenAction     `json:"actions,omitempty"`
	Links      []SirenLink       `json:"links,omitempty"`
}

func (e *SirenEmbed[P]) SetHref(href string) *SirenEmbed[P] {
	e.Href = href
	return e
}

func (e *SirenEmbed[P]) SetType(t string) *SirenEmbed[P] {
	e.Type = t
	return e
}

func (e *SirenEmbed[P]) SetTitle(title string) Siren[P] {
	e.Title = title
	return e
}

func (e *SirenEmbed[P]) SetClasses(classes []string) Siren[P] {
	e.Class = classes
	return e
}

func (e *SirenEmbed[P]) AddClass(class string) Siren[P] {
	if e.Class == nil {
		e.Class = []string{}
	}
	e.Class = append(e.Class, class)
	return e
}

func (e *SirenEmbed[P]) SetLinks(links []SirenLink) Siren[P] {
	e.Links = links
	return e
}

func (e *SirenEmbed[P]) AddLink(rel []string, href string, l SirenLink) Siren[P] {
	l.Rel = rel
	l.Href = href
	e.Links = append(e.Links, l)
	return e
}

func (e *SirenEmbed[P]) SetActions(actions []SirenAction) Siren[P] {
	e.Actions = actions
	return e
}

func (e *SirenEmbed[P]) AddAction(name string, title string, method string,
	href string, ctype string, fields ...SirenField) Siren[P] {
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

func (e *SirenEmbed[P]) SetEmbeds(embeds []SirenEmbed[any]) Siren[P] {
	e.Entities = embeds
	return e
}

// Rel is an argument because it is required
func (e *SirenEmbed[P]) AddEmbed(rel []string, embed SirenEmbed[any]) Siren[P] {
	embed.Rel = rel
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

var _ Siren[any] = (*SirenEmbed[any])(nil)
