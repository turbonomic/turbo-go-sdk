package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testProviderDTOGetProviderID(t *testing.T) {
	assert := assert.New(t)

	id := "should be random string"

	providerDto := &ProviderDTO{
		Id: &id,
	}

	assert.Equal(id, providerDto.getId())
}

func testProviderDTOGetProviderType(t *testing.T) {
	assert := assert.New(t)

	// TODO. Not sure if this is a good way to generate an EntityType. Or we hardcode one particular type here.
	pType := new(EntityDTO_EntityType)

	providerDto := &ProviderDTO{
		providerType: pType,
	}

	assert.Equal(pType, providerDto.getProviderType())
}
