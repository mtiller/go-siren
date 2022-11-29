package gosiren

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmbedJSON(t *testing.T) {
	assert := assert.New(t)

	siren := NewSirenEntity[any]().SetTitle("Welcome")
	siren.Title = "Welcome"

	str, err := json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"title":"Welcome"}`, string(str))

	child := NewSirenEntity[any]().SetTitle("I'm a child")

	siren.AddEmbed("#/child", child, "home")

	str, err = json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"title":"Welcome","entities":[{"rel":["home"],"href":"#/child","title":"I'm a child"}]}`, string(str))
}
