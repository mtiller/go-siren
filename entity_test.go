package gosiren

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntityJSON(t *testing.T) {
	assert := assert.New(t)

	siren := NewSirenEntity[any]()
	siren.Title = "Welcome"
	str, err := json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"title":"Welcome"}`, string(str))

	siren.AddLink([]string{"home"}, "https://example.com/home", SirenLink{
		Title: "Home",
	})

	str, err = json.Marshal(siren)
	assert.Nil(err)
	assert.Equal(`{"title":"Welcome","links":[{"rel":["home"],"href":"https://example.com/home","title":"Home"}]}`, string(str))
}
