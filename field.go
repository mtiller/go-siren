package gosiren

type SirenField struct {
	Name  string      `json:"name"`
	Class []string    `json:"class,omitempty"`
	Title string      `json:"title,omitempty"`
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

func MakeField(name string, vtype string, value string) SirenField {
	return SirenField{
		Name:  name,
		Type:  vtype,
		Value: value,
	}
}
