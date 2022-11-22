package gosiren

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

func (e *SirenEmbed) SetHref(href string) *SirenEmbed {
	e.Href = href
	return e
}

func (e *SirenEmbed) SetType(t string) *SirenEmbed {
	e.Type = t
	return e
}

func (e *SirenEmbed) SetTitle(title string) Siren {
	e.Title = title
	return e
}

func (e *SirenEmbed) SetClasses(classes []string) Siren {
	e.Class = classes
	return e
}

func (e *SirenEmbed) AddClass(class string) Siren {
	if e.Class == nil {
		e.Class = []string{}
	}
	e.Class = append(e.Class, class)
	return e
}

func (e *SirenEmbed) SetLinks(links []SirenLink) Siren {
	e.Links = links
	return e
}

func (e *SirenEmbed) AddLink(rel []string, href string, l SirenLink) Siren {
	l.Rel = rel
	l.Href = href
	e.Links = append(e.Links, l)
	return e
}

func (e *SirenEmbed) SetActions(actions []SirenAction) Siren {
	e.Actions = actions
	return e
}

func (e *SirenEmbed) AddAction(name string, title string, method string,
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

func (e *SirenEmbed) SetEmbeds(embeds []SirenEmbed) Siren {
	e.Entities = embeds
	return e
}

// Rel is an argument because it is required
func (e *SirenEmbed) AddEmbed(rel []string, embed SirenEmbed) Siren {
	embed.Rel = rel
	e.Entities = append(e.Entities, embed)
	return e
}

func NewSirenEmbed(rel []string) *SirenEmbed {
	return &SirenEmbed{
		Rel:        rel,
		Class:      []string{},
		Properties: map[string]interface{}{},
		Entities:   []SirenEmbed{},
		Actions:    []SirenAction{},
		Links:      []SirenLink{},
	}
}

var _ Siren = (*SirenEmbed)(nil)
