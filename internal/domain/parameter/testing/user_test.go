package testing

import (
	"testing"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/parameter"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestModifyPersonalProfileGetUpdateMap(t *testing.T) {
	t.Parallel()
	// Args
	type Args struct {
		Name  string
		Email string
	}
	var args Args
	_ = faker.FakeData(&args)
	mpp := parameter.ModifyPersonalProfile{
		Name:  args.Name,
		Email: args.Email,
	}
	updateMap := mpp.GetUpdateMap()
	assert.Equal(t, args.Name, updateMap["name"].(string))
	assert.Equal(t, args.Email, updateMap["email"].(string))
	mpp = parameter.ModifyPersonalProfile{}
	updateMap = mpp.GetUpdateMap()
	assert.Nil(t, updateMap["name"])
	assert.Nil(t, updateMap["email"])
}
