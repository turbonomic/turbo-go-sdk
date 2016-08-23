package supplychain

import (
	sdkproto "github.com/vmturbo/vmturbo-go-sdk/pkg/proto"
	"github.com/vmturbo/vmturbo-go-sdk/pkg/sdk"
)

var (
	cpuType sdkproto.CommodityDTO_CommodityType = sdkproto.CommodityDTO_CPU
	memType sdkproto.CommodityDTO_CommodityType = sdkproto.CommodityDTO_MEM

	//Commodity key is optional, when key is set, it serves as a constraint between seller and buyer
	//for example, the buyer can only go to a seller that sells the commodity with the required key
	cpuCommKey string = "cpu_comm_key"
	memCommKey string = "mem_comm_key"
)

type SupplyChainFactory struct{}

// SupplyChain definition: this function defines the buyer/seller relationships between each of
// the entity types in * the Target, the default Supply Chain definition in this function is:
// a Virtual Machine buyer, a Physical Machine seller and the commodities are CPU and Memory.
// Each entity type and the relationships are defined by a single TemplateDTO struct
// The function returns an array of TemplateDTO pointers
// TO MODIFY:
// For each entity: Create a supply chain builder object with sdk.NewSupplyChainNodeBuilder()
//		    Set a provider type if the new entity is a buyer , create commodity objects
//		    and add them to the entity's supply chain builder object
//                  Add commodity objects with the selling function to the entity you create if
//		    it is a seller.
//		    Add the new entity to the supplyChainBuilder instance with either the Top()
//		    or  Entity() methods
// The SupplyChainBuilder() function is only called once, in this function.
func (this *SupplyChainFactory) CreateSupplyChain() []*sdkproto.TemplateDTO {
	vmSupplyChainNodeBuilder := this.virtualMachineSupplyChainBuilder()
	pmSupplyChainNodeBuilder := this.physicalMachineSupplyChainBuilder()

	// SupplyChain building
	// The last buyer in the supply chain is set as the top entity with the Top() method
	// All other entities are added to the SupplyChainBuilder with the Entity() method
	supplyChainBuilder := sdk.NewSupplyChainBuilder()
	supplyChainBuilder.Top(vmSupplyChainNodeBuilder)
	supplyChainBuilder.Entity(pmSupplyChainNodeBuilder)

	return supplyChainBuilder.Create()
}

// Create supply chain definition for Physical Machine.
func (this *SupplyChainFactory) physicalMachineSupplyChainBuilder() *sdk.SupplyChainBuilder {
	// PM Creation Process
	pmSupplyChainNodeBuilder := sdk.NewSupplyChainNodeBuilder()
	// Creates a Physical Machine entity and sets the type of commodity it sells to CPU
	pmSupplyChainNodeBuilder = pmSupplyChainNodeBuilder.
		Entity(sdkproto.EntityDTO_PHYSICAL_MACHINE).
		Selling(sdkproto.CommodityDTO_CPU, cpuCommKey).
		Selling(sdkproto.CommodityDTO_MEM, memCommKey)

	return pmSupplyChainNodeBuilder
}

// Create supply chain definition for Vitual Machine
func (this *SupplyChainFactory) virtualMachineSupplyChainBuilder() *sdk.SupplyChainBuilder {
	// VM Creation Process
	vmSupplyChainNodeBuilder := sdk.NewSupplyChainNodeBuilder()
	// Creates a Virtual Machine entity
	vmSupplyChainNodeBuilder = vmSupplyChainNodeBuilder.Entity(sdkproto.EntityDTO_VIRTUAL_MACHINE)
	cpuTemplateComm := &sdkproto.TemplateCommodity{
		Key:           &cpuCommKey,
		CommodityType: &cpuType,
	}
	memTemplateComm := &sdk.TemplateCommodity{
		Key:           &memCommKey,
		CommodityType: &memType,
	}
	// The Entity type for the Virtual Machine's commodity provider is defined by the Provider() method.
	// The Commodity type for Virtual Machine's buying relationship is define by the Buys() method
	vmSupplyChainNodeBuilder = vmSupplyChainNodeBuilder.
		Provider(sdk.EntityDTO_PHYSICAL_MACHINE, sdk.Provider_HOSTING).
		Buys(*cpuTemplateComm).
		Buys(*memTemplateComm)

	return vmSupplyChainNodeBuilder
}
