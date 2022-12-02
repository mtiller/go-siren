package siren

type Embeddable interface {
	GetTitle() string
	GetClass() []string
	GetProperties() interface{}
	GetActions() []SirenAction
	GetLinks() []SirenLink
	GetEntities() []EmbeddedEntity[any]
}
