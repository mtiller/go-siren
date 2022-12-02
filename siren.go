package siren

import (
	"encoding/json"
)

const MimeType = "application/vnd.siren+json"

type Siren[P any] struct {
	Title      string                `json:"title,omitempty"`
	Class      []string              `json:"class,omitempty"`
	Properties P                     `json:"properties,omitempty"`
	Actions    []SirenAction         `json:"actions,omitempty"`
	Links      []SirenLink           `json:"links,omitempty"`
	Entities   []EmbeddedEntity[any] `json:"entities,omitempty"`
}

func (e *Siren[P]) GetTitle() string {
	return e.Title
}

func (e *Siren[P]) SetTitle(title string) *Siren[P] {
	e.Title = title
	return e
}

func (e *Siren[P]) GetClass() []string {
	return e.Class
}

func (e *Siren[P]) SetClasses(classes []string) *Siren[P] {
	e.Class = classes
	return e
}

func (e *Siren[P]) AddClass(class string) *Siren[P] {
	if e.Class == nil {
		e.Class = []string{}
	}
	e.Class = append(e.Class, class)
	return e
}

func (e *Siren[P]) GetProperties() interface{} {
	return e.Properties
}

func (e *Siren[P]) SetProperties(props P) *Siren[P] {
	e.Properties = props
	return e
}

func (e *Siren[P]) GetLinks() []SirenLink {
	return e.Links
}

func (e *Siren[P]) SetLinks(links []SirenLink) *Siren[P] {
	e.Links = links
	return e
}

func (e *Siren[P]) AddLink(href string, l SirenLink, r1 string, rel ...string) *Siren[P] {
	l.Rel = append([]string{r1}, rel...)
	l.Href = href
	e.Links = append(e.Links, l)
	return e
}

func (e *Siren[P]) GetActions() []SirenAction {
	return e.Actions
}

func (e *Siren[P]) SetActions(actions []SirenAction) *Siren[P] {
	e.Actions = actions
	return e
}

func (e *Siren[P]) AddAction(name string, title string, method string,
	href string, ctype string, fields ...SirenField) *Siren[P] {
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

func (e *Siren[P]) GetEntities() []EmbeddedEntity[any] {
	return e.Entities
}

func (e *Siren[P]) AddEntity(href string, siren Embeddable, r1 string, rel ...string) *Siren[P] {
	embed := EmbeddedEntity[any]{
		Rel:        append([]string{r1}, rel...),
		Href:       href,
		Title:      siren.GetTitle(),
		Class:      siren.GetClass(),
		Properties: siren.GetProperties(),
		Entities:   siren.GetEntities(),
		Actions:    siren.GetActions(),
		Links:      siren.GetLinks(),
	}

	e.Entities = append(e.Entities, embed)
	return e
}

func (e *Siren[P]) LinksWithClass(c string) []SirenLink {
	if e.Links == nil {
		return []SirenLink{}
	}
	return Filter(e.Links, func(link SirenLink) bool {
		if link.Class != nil {
			for _, cl := range link.Class {
				if cl == c {
					return true
				}
			}
		}
		return false
	})
}

func (e *Siren[P]) LinksWithRel(r string) []SirenLink {
	if e.Links == nil {
		return []SirenLink{}
	}
	return Filter(e.Links, func(link SirenLink) bool {
		if link.Rel != nil {
			for _, rl := range link.Rel {
				if rl == r {
					return true
				}
			}
		}
		return false
	})
}

func (e *Siren[P]) EntitiesWithClass(c string) []EmbeddedEntity[any] {
	if e.Entities == nil {
		return []EmbeddedEntity[any]{}
	}
	return Filter(e.Entities, func(entity EmbeddedEntity[any]) bool {
		if entity.Class != nil {
			for _, cl := range entity.Class {
				if cl == c {
					return true
				}
			}
		}
		return false
	})
}

func (e *Siren[P]) EntitiesWithRel(r string) []EmbeddedEntity[any] {
	if e.Entities == nil {
		return []EmbeddedEntity[any]{}
	}
	return Filter(e.Entities, func(entity EmbeddedEntity[any]) bool {
		if entity.Rel != nil {
			for _, rl := range entity.Rel {
				if rl == r {
					return true
				}
			}
		}
		return false
	})
}

func NewSiren() *Siren[any] {
	return &Siren[any]{
		Class:    []string{},
		Entities: []EmbeddedEntity[any]{},
		Actions:  []SirenAction{},
		Links:    []SirenLink{},
	}
}

func NewSirenP[P any](p P) *Siren[P] {
	return &Siren[P]{
		Class:      []string{},
		Properties: p,
		Entities:   []EmbeddedEntity[any]{},
		Actions:    []SirenAction{},
		Links:      []SirenLink{},
	}
}

func Parse(s string) (*Siren[any], error) {
	ret := Siren[any]{}
	err := json.Unmarshal([]byte(s), &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

var _ Embeddable = (*Siren[any])(nil)
var _ Embeddable = (*Siren[struct{}])(nil)
