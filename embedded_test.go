package siren

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Prop1 struct {
	X int
}

func TestEmbedJSON(t *testing.T) {
	assert := assert.New(t)

	siren := NewSiren().SetTitle("Welcome")

	str, err := json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"title":"Welcome"}`, string(str))

	child := NewSirenP(Prop1{X: 5}).SetTitle("I'm a child")

	siren.AddEntity(child, "home")
	siren.Entities[0].AddClass("foo")

	str, err = json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"title":"Welcome","entities":[{"rel":["home"],"href":"#/child","title":"I'm a child","class":["foo"],"properties":{"X":5}}]}`, string(str))

	siren2, err := Parse(string(str))
	assert.Nil(err)
	str2, err := json.Marshal(siren2)
	assert.Nil(err)
	assert.Equal(str, str2)
}

func TestContruction(t *testing.T) {
	assert := assert.New(t)

	siren := NewSiren().
		SetTitle("Welcome").
		AddEntity(NewSiren().
			SetTitle("Child").
			AddEntity(NewSiren().
				SetTitle("Grandchild"), "item"),
			"item")
	str, err := json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"title":"Welcome","entities":[{"rel":["item"],"href":"#/child","title":"Child","entities":[{"rel":["item"],"href":"#/grandchild","title":"Grandchild"}]}]}`, string(str))

	str, err = json.MarshalIndent(siren, "", "  ")
	assert.Nil(err)
	assert.Equal(`{
  "title": "Welcome",
  "entities": [
    {
      "rel": [
        "item"
      ],
      "href": "#/child",
      "title": "Child",
      "entities": [
        {
          "rel": [
            "item"
          ],
          "href": "#/grandchild",
          "title": "Grandchild"
        }
      ]
    }
  ]
}`, string(str))
}
