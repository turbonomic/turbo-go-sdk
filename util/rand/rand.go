package rand

import (
	"math/rand"
	"sync"
	"time"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"math"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
var numLetters = len(letters)
var rng = struct {
	sync.Mutex
	rand *rand.Rand
}{
	rand: rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
}

// String generates a random alphanumeric string n characters long.  This will
// panic if n is less than zero.
func String(n int) string {
	if n < 0 {
		panic("out-of-bounds value")
	}
	b := make([]rune, n)
	rng.Lock()
	defer rng.Unlock()
	for i := range b {
		b[i] = letters[rng.rand.Intn(numLetters)]
	}
	return string(b)
}

// Seed seeds the rng with the provided seed.
func Seed(seed int64) {
	rng.Lock()
	defer rng.Unlock()

	rng.rand = rand.New(rand.NewSource(seed))
}

// Create a random entity type.
func RandomEntityType() proto.EntityDTO_EntityType {
	return proto.EntityDTO_EntityType(rand.Int31n(42))
}

// Create a random commodity type.
func RandomCommodityType() proto.CommodityDTO_CommodityType {
	return proto.CommodityDTO_CommodityType(rand.Int31n(77))
}

// Create a random power state value, range from 1 to 4.
func RandomPowerState() proto.EntityDTO_PowerState {
	return proto.EntityDTO_PowerState(rand.Int31n(4) + 1)
}

// Create a random entity origin, range from 1 to 2.
func RandomOrigin() proto.EntityDTO_EntityOrigin {
	return proto.EntityDTO_EntityOrigin(rand.Int31n(2) + 1)
}

// Create a random commodityDTO bought.
func RandomCommodityDTOBought() *proto.CommodityDTO {
	// a random commodity type.
	cType := RandomCommodityType()
	// a random key
	key := String(5)
	// a random used
	used := rand.Float64()
	return &proto.CommodityDTO{
		CommodityType: &cType,
		Key:           &key,
		Used:          &used,
	}

}

// Create a random CommodityDTO sold.
func RandomCommodityDTOSold() *proto.CommodityDTO {
	// a random commodity type.
	cType := RandomCommodityType()
	// a random key
	key := String(5)
	// a random capacity
	capacity := rand.Float64()
	// a random used
	used := rand.Float64()
	return &proto.CommodityDTO{
		CommodityType: &cType,
		Key:           &key,
		Capacity:      &capacity,
		Used:          &used,
	}

}

func RandomExternalEntityLink_ServerEntityPropDef() *proto.ExternalEntityLink_ServerEntityPropDef {
	entity := RandomEntityType()
	attribute := String(5)
	return &proto.ExternalEntityLink_ServerEntityPropDef{
		Entity:    &entity,
		Attribute: &attribute,
	}
}

func RandomTemplateCommodity() *proto.TemplateCommodity {
	// a random commodity type.
	cType := RandomCommodityType()
	// a random key
	key := String(5)
	return &proto.TemplateCommodity{
		CommodityType: &cType,
		Key:           &key,
	}
}

func RandomProvider() *proto.Provider {
	providerEntityType := RandomEntityType()
	relationShip := RandomProviderConsumerRelationship()
	maxCardinality := int32(math.MaxInt32)
	minCardinality := int32(0)
	return &proto.Provider{
		TemplateClass:  &providerEntityType,
		ProviderType:   &relationShip,
		CardinalityMax: &maxCardinality,
		CardinalityMin: &minCardinality,
	}
}

func RandomProviderConsumerRelationship() proto.Provider_ProviderType {
	return proto.Provider_ProviderType(rand.Int31n(2))
}

func RandomApplicationData() *proto.EntityDTO_ApplicationData {
	t := String(10)
	i := String(10)
	return &proto.EntityDTO_ApplicationData{
		Type:      &t,
		IpAddress: &i,
	}
}

func RandomVirtualMachineData() *proto.EntityDTO_VirtualMachineData {
	ipAddress := []string{String(14)}
	gName := String(5)
	return &proto.EntityDTO_VirtualMachineData{
		IpAddress:      ipAddress,
		VmState:        RandomVMState(),
		GuestName:      &gName,
		AnnotationNote: []*proto.EntityDTO_VirtualMachineData_AnnotationNote{RandomAnnotationNote()},
	}
}

func RandomVirtualApplicationData() *proto.EntityDTO_VirtualApplicationData {
	ipAddress := String(14)
	serviceType := String(5)
	t := String(5)
	port := rand.Int31n(9999)
	return &proto.EntityDTO_VirtualApplicationData{
		Type:        &t,
		Port:        &port,
		IpAddress:   &ipAddress,
		ServiceType: &serviceType,
	}
}

func RandomVMState() *proto.EntityDTO_VMState {
	connected := rand.Int31n(2) == 1
	return &proto.EntityDTO_VMState{
		Connected: &connected,
	}
}

func RandomAnnotationNote() *proto.EntityDTO_VirtualMachineData_AnnotationNote {
	key := String(5)
	value := String(5)
	return &proto.EntityDTO_VirtualMachineData_AnnotationNote{
		Key:   &key,
		Value: &value,
	}
}

func RandomAccoutValue() *proto.AccountValue {
	key := String(5)
	value := String(5)
	return &proto.AccountValue{
		Key:         &key,
		StringValue: &value,
	}
}
