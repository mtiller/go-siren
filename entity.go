package gosiren

type SirenEntity[P any] struct {
	Title      string            `json:"title,omitempty"`
	Class      []string          `json:"class,omitempty"`
	Properties P                 `json:"properties,omitempty"`
	Entities   []SirenEmbed[any] `json:"entities,omitempty"`
	Actions    []SirenAction     `json:"actions,omitempty"`
	Links      []SirenLink       `json:"links,omitempty"`
}

func (e *SirenEntity[P]) SetTitle(title string) Siren[P] {
	e.Title = title
	return e
}

func (e *SirenEntity[P]) SetClasses(classes []string) Siren[P] {
	e.Class = classes
	return e
}

func (e *SirenEntity[P]) AddClass(class string) Siren[P] {
	if e.Class == nil {
		e.Class = []string{}
	}
	e.Class = append(e.Class, class)
	return e
}

func (e *SirenEntity[P]) SetLinks(links []SirenLink) Siren[P] {
	e.Links = links
	return e
}

func (e *SirenEntity[P]) AddLink(rel []string, href string, l SirenLink) Siren[P] {
	l.Rel = rel
	l.Href = href
	e.Links = append(e.Links, l)
	return e
}

func (e *SirenEntity[P]) SetActions(actions []SirenAction) Siren[P] {
	e.Actions = actions
	return e
}

func (e *SirenEntity[P]) AddAction(name string, title string, method string,
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

func (e *SirenEntity[P]) SetEmbeds(embeds []SirenEmbed[any]) Siren[P] {
	e.Entities = embeds
	return e
}

// Rel is an argument because it is required
func (e *SirenEntity[P]) AddEmbed(rel []string, embed SirenEmbed[any]) Siren[P] {
	embed.Rel = rel
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

var _ Siren[any] = (*SirenEntity[any])(nil)
