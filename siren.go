package gosiren

type Siren interface {
	SetTitle(string) Siren
	SetClasses([]string) Siren
	AddClass(string) Siren
	// SetProperties(map[string]interface{}) Siren
	SetLinks([]SirenLink) Siren
	AddLink([]string, string, SirenLink) Siren
	SetActions([]SirenAction) Siren
	AddAction(string, string, string, string, string, ...SirenField) Siren
	SetEmbeds([]SirenEmbed) Siren
	AddEmbed([]string, SirenEmbed) Siren
}
