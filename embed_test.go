package gosiren

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmbedJSON(t *testing.T) {
	assert := assert.New(t)

	siren := NewSirenEmbed[any]([]string{"next"})
	siren.Title = "Welcome"
	str, err := json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"rel":["next"],"title":"Welcome"}`, string(str))

	child := NewSirenEmbed[struct{}]([]string{"next"})
	child.SetTitle("Home")

	siren.AddEmbed([]string{"home"}, child)

	str, err = json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"rel":["next"],"title":"Welcome","entities":[{"rel":["home"],"title":"Home"}]}`, string(str))
}
