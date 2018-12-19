package group

import "github.com/turbonomic/turbo-go-sdk/pkg/proto"

type groupData struct {
	matchingBuyers *Matching
	entities       []string
	entityTypePtr  *proto.EntityDTO_EntityType
}

type BuyerPolicyData struct {
	*groupData
	atMost int32
}

type SellerPolicyData struct {
	*groupData
}

func StaticBuyers(buyers []string) *BuyerPolicyData {
	buyerData := &BuyerPolicyData{
		groupData: &groupData{entities: buyers},
	}
	return buyerData
}

func DynamicBuyers(matchingBuyers *Matching) *BuyerPolicyData {
	buyerData := &BuyerPolicyData{
		groupData: &groupData{matchingBuyers: matchingBuyers},
	}
	return buyerData
}

func (buyer *BuyerPolicyData) OfType(buyerType proto.EntityDTO_EntityType) *BuyerPolicyData {
	buyer.entityTypePtr = &buyerType
	return buyer
}

func (buyer *BuyerPolicyData) AtMost(maxBuyers int32) *BuyerPolicyData {
	buyer.atMost = maxBuyers
	return buyer
}

func StaticSellers(buyers []string) *SellerPolicyData {
	buyerData := &SellerPolicyData{groupData: &groupData{}}
	buyerData.entities = buyers
	return buyerData
}

func DynamicSellers(matchingBuyers *Matching) *SellerPolicyData {
	buyerData := &SellerPolicyData{groupData: &groupData{}}
	buyerData.matchingBuyers = matchingBuyers
	return buyerData
}

func (buyer *SellerPolicyData) OfType(buyerType proto.EntityDTO_EntityType) *SellerPolicyData {
	buyer.entityTypePtr = &buyerType
	return buyer
}
