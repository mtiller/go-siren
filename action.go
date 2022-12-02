package siren

type SirenAction struct {
	Name   string       `json:"name"`
	Href   string       `json:"href"`
	Class  []string     `json:"class,omitempty"`
	Method string       `json:"method,omitempty"`
	Title  string       `json:"title,omitempty"`
	Type   string       `json:"type,omitempty"`
	Fields []SirenField `json:"fields,omitempty"`
}

func (a *SirenAction) SetTitle(title string) *SirenAction {
	a.Title = title
	return a
}

func (a *SirenAction) SetClasses(classes []string) *SirenAction {
	a.Class = classes
	return a
}

func (a *SirenAction) AddClass(class string) *SirenAction {
	if a.Class == nil {
		a.Class = []string{}
	}
	a.Class = append(a.Class, class)
	return a
}

func (a *SirenAction) SetFields(fields []SirenField) *SirenAction {
	a.Fields = fields
	return a
}

func (a *SirenAction) AddField(field SirenField) *SirenAction {
	if a.Fields == nil {
		a.Fields = []SirenField{}
	}
	a.Fields = append(a.Fields, field)
	return a
}

func (a *SirenAction) SetMethod(method string) *SirenAction {
	a.Method = method
	return a
}

func (a *SirenAction) SetType(t string) *SirenAction {
	a.Type = t
	return a
}

func MakeAction(name string, href string) SirenAction {
	return SirenAction{
		Name:   name,
		Href:   href,
		Class:  []string{},
		Fields: []SirenField{},
	}
}
