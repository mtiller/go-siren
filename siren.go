package gosiren

type Siren[P any] interface {
	SetTitle(string) Siren[P]
	SetClasses([]string) Siren[P]
	AddClass(string) Siren[P]
	// SetProperties(map[string]interface{}) Siren
	SetLinks([]SirenLink) Siren[P]
	AddLink([]string, string, SirenLink) Siren[P]
	SetActions([]SirenAction) Siren[P]
	AddAction(string, string, string, string, string, ...SirenField) Siren[P]
	SetEmbeds([]SirenEmbed[any]) Siren[P]
	AddEmbed([]string, SirenEmbed[any]) Siren[P]
}
