package gosiren

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyProperties struct {
	X string `json:"x"`
	Y int    `json:"y"`
}

func TestEntityJSON(t *testing.T) {
	assert := assert.New(t)

	siren := NewSirenEntity[any]().SetTitle("Welcome")

	str, err := json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"title":"Welcome"}`, string(str))

	siren.AddLink([]string{"home"}, "https://example.com/home", SirenLink{
		Title: "Home",
	})

	str, err = json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"title":"Welcome","links":[{"rel":["home"],"href":"https://example.com/home","title":"Home"}]}`, string(str))

	mysiren := NewSirenEntity[MyProperties]()
	mysiren.Title = "Uses MyProperties"
	mysiren.SetProperties(&MyProperties{X: "hello", Y: 5})

	str, err = json.Marshal(mysiren)
	assert.Nil(err)
	assert.Equal(`{"title":"Uses MyProperties","properties":{"x":"hello","y":5}}`, string(str))
}
