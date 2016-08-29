package common

import "github.com/vmturbo/vmturbo-go-sdk/pkg/proto"

type ProviderDTO struct {
	providerType *proto.EntityDTO_EntityType
	Id           *string
}

func CreateProvider(pType proto.EntityDTO_EntityType, id string) *ProviderDTO {
	return &ProviderDTO{
		providerType: &pType,
		Id:           &id,
	}
}

func (pDto *ProviderDTO) getProviderType() *proto.EntityDTO_EntityType {
	return pDto.providerType
}

func (pDto *ProviderDTO) getId() *string {
	return pDto.Id
}
