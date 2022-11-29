package gosiren

type Embeddable interface {
	GetTitle() string
	GetClass() []string
	GetProperties() interface{}
	GetActions() []SirenAction
	GetLinks() []SirenLink
	GetEntities() []SirenEmbed[any]
}
