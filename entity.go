package gosiren

type SirenEntity struct {
	Title      string                 `json:"title,omitempty"`
	Class      []string               `json:"class,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Entities   []SirenEmbed           `json:"entities,omitempty"`
	Actions    []SirenAction          `json:"actions,omitempty"`
	Links      []SirenLink            `json:"links,omitempty"`
}

func (e *SirenEntity) SetTitle(title string) Siren {
	e.Title = title
	return e
}

func (e *SirenEntity) SetClasses(classes []string) Siren {
	e.Class = classes
	return e
}

func (e *SirenEntity) AddClass(class string) Siren {
	if e.Class == nil {
		e.Class = []string{}
	}
	e.Class = append(e.Class, class)
	return e
}

func (e *SirenEntity) SetLinks(links []SirenLink) Siren {
	e.Links = links
	return e
}

func (e *SirenEntity) AddLink(rel []string, href string, l SirenLink) Siren {
	l.Rel = rel
	l.Href = href
	e.Links = append(e.Links, l)
	return e
}

func (e *SirenEntity) SetActions(actions []SirenAction) Siren {
	e.Actions = actions
	return e
}

func (e *SirenEntity) AddAction(name string, title string, method string,
	href string, ctype string, fields ...SirenField) Siren {
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

func (e *SirenEntity) SetEmbeds(embeds []SirenEmbed) Siren {
	e.Entities = embeds
	return e
}

// Rel is an argument because it is required
func (e *SirenEntity) AddEmbed(rel []string, embed SirenEmbed) Siren {
	embed.Rel = rel
	e.Entities = append(e.Entities, embed)
	return e
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

var _ Siren = (*SirenEntity)(nil)
