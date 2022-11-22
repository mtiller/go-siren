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

	siren.AddLink([]string{"home"}, "https://example.com/home", SirenLink{
		Title: "Home",
	})

	str, err = json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"rel":["next"],"title":"Welcome","links":[{"rel":["home"],"href":"https://example.com/home","title":"Home"}]}`, string(str))
}
