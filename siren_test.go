package siren

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

	siren := NewSiren().SetTitle("Welcome").SetProperties(nil)

	str, err := json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"title":"Welcome"}`, string(str))

	siren.AddLink("https://example.com/home", SirenLink{
		Title: "Home",
	}, "home")

	str, err = json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"title":"Welcome","links":[{"rel":["home"],"href":"https://example.com/home","title":"Home"}]}`, string(str))

	mysiren := NewSirenP(&MyProperties{X: "hello", Y: 5})
	mysiren.Title = "Uses MyProperties"

	str, err = json.Marshal(mysiren)
	assert.Nil(err)
	assert.Equal(`{"title":"Uses MyProperties","properties":{"x":"hello","y":5}}`, string(str))
}

func TestQuery(t *testing.T) {
	assert := assert.New(t)

	siren := NewSiren().SetTitle("Welcome")
	assert.Equal(0, len(siren.Links))

	foos := siren.LinksWithClass("foo")
	assert.Equal(0, len(foos))

	siren.AddLink("#/child/foo", SirenLink{Class: []string{"foo"}}, "item")
	siren.AddLink("#/child/bar", SirenLink{Class: []string{"bar"}}, "item", "collection")

	foos = siren.LinksWithClass("foo")
	assert.Equal(1, len(foos))

	siren.AddEntity("#/child/foo", NewSiren().AddClass("foo"), "item")
	siren.AddEntity("#/child/bar", NewSiren().AddClass("bar"), "item", "collection")

	fooes := siren.EntitiesWithClass("foo")
	assert.Equal(1, len(fooes))

	assert.Equal(2, len(siren.LinksWithRel("item")))
	assert.Equal(1, len(siren.LinksWithRel("collection")))

	assert.Equal(2, len(siren.EntitiesWithRel("item")))
	assert.Equal(1, len(siren.EntitiesWithRel("collection")))
}
