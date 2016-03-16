package main

import (
	"github.com/golang/glog"
	"github.com/vmturbo/vmturbo-go-sdk/communicator"
	"github.com/vmturbo/vmturbo-go-sdk/sdk"
)

var wsCommunicator *communicator.WebSocketCommunicator
var loginInfo *ConnectionInfo
var vMTApiRequestHandler *VMTApiRequestHandler
var msgHandler *MsgHandler

// Struct which holds identifying information for connecting to the VMTServer
type ConnectionInfo struct {
	ServerAddr         string
	LocalAddr          string
	Type               string
	OpsManagerUsername string
	OpsManagerPassword string
	// The user defined name is for UI admin display purposes only, this field must be set by the user with a name or the IP address of the target
	UserDefinedNameorIPAddress string
	Username                   string
	Password                   string
	TargetIdentifier           string
}

// A function that creates an array of *sdk.CommodityDTO , this array defines all the commodities bought by a single
// entity in the target supply chain.
// returns an array of *sdk.CommodityDTO
func CreateCommoditiesBought(comms_array []*Commodity_Params) []*sdk.CommodityDTO {
	var commoditiesBought []*sdk.CommodityDTO
	for _, comm := range comms_array {
		cpuComm := sdk.NewCommodtiyDTOBuilder(comm.commType).Key(comm.commKey).Capacity(comm.cap).Used(comm.used).Create()
		commoditiesBought = append(commoditiesBought, cpuComm)
	}
	return commoditiesBought
}

// A function that creates an array of *sdk.CommodityDTO , this array defines all the commodities sold by a single
// entity
func CreateCommoditiesSold(comms_array []*Commodity_Params) []*sdk.CommodityDTO {
	var commoditiesSold []*sdk.CommodityDTO
	for _, comm := range comms_array {
		cpuComm := sdk.NewCommodtiyDTOBuilder(comm.commType).Key(comm.commKey).Capacity(comm.cap).Used(comm.used).Create()
		commoditiesSold = append(commoditiesSold, cpuComm)
	}
	return commoditiesSold
}

// This function builds an EntityDTO object from information provided from the one of the entities at discovery time.
// It returns an *sdk.EntityDTO which points to the EntityDTO created.
func (e *Entity_Params) buildEntityDTO() *sdk.EntityDTO {
	entityDTOBuilder := sdk.NewEntityDTOBuilder(e.entityType, e.entityID)
	entityDTOBuilder.DisplayName(e.entityDisplayName)
	if e.isBuyer == true {
		entityDTOBuilder.SetProvider(e.providerType, e.providerID)
		commoditiesbought := CreateCommoditiesBought(e.commoditiesBought)
		entityDTOBuilder = entityDTOBuilder.BuysCommodities(commoditiesbought)
		glog.Infof("after buys --> %++v", entityDTOBuilder)
	}
	if e.isSeller == true {
		commoditiesSold := CreateCommoditiesSold(e.commoditiesSold)
		for _, curcomm := range commoditiesSold {
			entityDTOBuilder = entityDTOBuilder.Sells(*curcomm.CommodityType, *curcomm.Key).Capacity(*curcomm.Capacity).Used(*curcomm.Used)
		}
	}
	entityDTO := entityDTOBuilder.Create()
	glog.Infof("after create entityDTO %d", len(entityDTO.CommoditiesBought))
	return entityDTO
}

// A struct that contains a Node probe, there is one Kubernetes probe and one NodeProbe for each target
type TargetProbe struct {
	//	nodeProbe *NodeProbe
	entities []*Entity_Params
}

// Struct that holds parameters for each commodity sold or bought by a given entity
type Commodity_Params struct {
	commType sdk.CommodityDTO_CommodityType
	commKey  string
	used     float64
	cap      float64
}

// Struct that holds a given entity's identifying and property information
type Entity_Params struct {
	isBuyer           bool
	isSeller          bool
	entityType        sdk.EntityDTO_EntityType
	entityID          string
	entityDisplayName string
	commoditiesSold   []*Commodity_Params
	commoditiesBought []*Commodity_Params
	providerID        string
	providerType      sdk.EntityDTO_EntityType
}

// Method that creates an array of entities found at this target
// Creates arrays of bought/sold commodities for each entity
// Sets the entities field of the NodeProbe it is called on to the
// newly created entity array
// Default: array of sold commodities contains a sdk.CommodityDTO_CPU
//	    and a sdk.CommodityDTO_MEM
//	    array of bought commodities contains a sdk.CommodityDTO_CPU
// 	    and asdk.CommodityDTO_MEM
//          Array of Entities contains 2 sellers, selling the same commodities
//          and 3 buyers, two buyers buying from seller 1 and one buyer buying from
//          seller 2.
func (targetProbe *TargetProbe) SampleProbe() {
	var s_comms_array []*Commodity_Params
	var b_comms_array []*Commodity_Params
	var entities []*Entity_Params
	comm1 := &Commodity_Params{
		commType: sdk.CommodityDTO_CPU,
		commKey:  "cpu_comm",
		used:     4,
		cap:      100,
	}
	comm2 := &Commodity_Params{
		commType: sdk.CommodityDTO_MEM,
		commKey:  "mem_comm",
		used:     10,
		cap:      100,
	}
	comm3 := &Commodity_Params{
		commType: sdk.CommodityDTO_CPU,
		commKey:  "cpu_comm",
		used:     0,
		cap:      1000,
	}
	comm4 := &Commodity_Params{
		commType: sdk.CommodityDTO_MEM,
		commKey:  "mem_comm",
		used:     0,
		cap:      1000,
	}
	s_comms_array = append(s_comms_array, comm1)
	s_comms_array = append(s_comms_array, comm2)
	b_comms_array = append(b_comms_array, comm3)
	b_comms_array = append(b_comms_array, comm4)

	newSeller1 := &Entity_Params{
		isBuyer:           false,
		isSeller:          true,
		entityType:        sdk.EntityDTO_PHYSICAL_MACHINE,
		entityID:          "PM_seller_1",
		entityDisplayName: "test_PM_seller_1",
		commoditiesSold:   s_comms_array,
	}
	newSeller2 := &Entity_Params{
		isBuyer:           false,
		isSeller:          true,
		entityType:        sdk.EntityDTO_PHYSICAL_MACHINE,
		entityID:          "PM_seller_2",
		entityDisplayName: "test_PM_seller_2",
		commoditiesSold:   s_comms_array,
	}
	newBuyer1 := &Entity_Params{
		isBuyer:           true,
		isSeller:          false,
		entityType:        sdk.EntityDTO_VIRTUAL_MACHINE,
		entityID:          "VM_buyer_1A",
		entityDisplayName: "test_VM_buyer_1A",
		commoditiesBought: b_comms_array,
		providerID:        "PM_seller_1",
	}
	newBuyer2 := *newBuyer1
	newBuyer2.entityID = "VM_buyer_1B"
	newBuyer2.entityDisplayName = "test_VM_buyer_1B"
	newBuyer3 := *newBuyer1
	newBuyer3.entityID = "VM_buyer_2A"
	newBuyer3.entityDisplayName = "test_VM_buyer_2A"
	newBuyer3.providerID = "PM_seller_2"
	entities = append(entities, newSeller1)
	entities = append(entities, newSeller2)
	entities = append(entities, newBuyer1)
	entities = append(entities, &newBuyer2)
	entities = append(entities, &newBuyer3)
	targetProbe.entities = entities
}

// this function turns our NodeArray from the Kubernetes.NodeProbe as a []*sdk.EntityDTO
func (probe *TargetProbe) getNodeEntityDTOs() []*sdk.EntityDTO {
	probe.SampleProbe()
	// create PM or VM EntityDTO
	var entityDTOarray []*sdk.EntityDTO
	for _, entity := range probe.entities {
		// we call createCommoditySold to get []*sdk.CommodityDTO
		newEntityDTO := entity.buildEntityDTO()
		entityDTOarray = append(entityDTOarray, newEntityDTO)
	}

	return entityDTOarray
}

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
func createSupplyChain() []*sdk.TemplateDTO {
	//Commodity key is optional, when key is set, it serves as a constraint between seller and buyer
	//for example, the buyer can only go to a seller that sells the commodity with the required key
	optionalKey := "commodity_key"
	// VM Creation Process
	vmSupplyChainNodeBuilder := sdk.NewSupplyChainNodeBuilder()
	// Creates a Virtual Machine entity
	vmSupplyChainNodeBuilder = vmSupplyChainNodeBuilder.Entity(sdk.EntityDTO_VIRTUAL_MACHINE)
	cpuType := sdk.CommodityDTO_CPU
	cpuTemplateComm := &sdk.TemplateCommodity{
		Key:           &optionalKey,
		CommodityType: &cpuType,
	}

	memType := sdk.CommodityDTO_MEM
	memTemplateComm := &sdk.TemplateCommodity{
		Key:           &optionalKey,
		CommodityType: &memType,
	}
	// The Entity type for the Virtual Machine's commodity provider is defined by the Provider() method.
	// The Commodity type for Virtual Machine's buying relationship is define by the Buys() method
	vmSupplyChainNodeBuilder = vmSupplyChainNodeBuilder.Provider(sdk.EntityDTO_PHYSICAL_MACHINE, sdk.Provider_HOSTING).Buys(*cpuTemplateComm).Buys(*memTemplateComm)

	// PM Creation Process
	pmSupplyChainNodeBuilder := sdk.NewSupplyChainNodeBuilder()
	// Creates a Physical Machine entity and sets the type of commodity it sells to CPU
	pmSupplyChainNodeBuilder = pmSupplyChainNodeBuilder.Entity(sdk.EntityDTO_PHYSICAL_MACHINE).Selling(sdk.CommodityDTO_CPU, optionalKey).Selling(sdk.CommodityDTO_MEM, optionalKey)
	// SupplyChain building
	//  The last buyer in the supply chain is set as the top entity with the Top() method
	// All other entities are added to the SupplyChainBuilder with the Entity() method
	supplyChainBuilder := sdk.NewSupplyChainBuilder()
	supplyChainBuilder.Top(vmSupplyChainNodeBuilder)
	supplyChainBuilder.Entity(pmSupplyChainNodeBuilder)

	return supplyChainBuilder.Create()
}

func init() {
	//
	//User defined settings
	//
	local_IP := "172.16.162.149"
	VMTServer_IP := "160.39.162.190"
	TargetIdentifier := "userDefinedTarget"
	OpsManagerUsername := "administrator"
	OpsManagerPassword := "a"
	localAddress := "ws://" + local_IP
	VMTServerAddress := VMTServer_IP + ":8080"
	wsCommunicator = new(communicator.WebSocketCommunicator)
	wsCommunicator.ServerUsername = "vmtRemoteMediation"
	wsCommunicator.ServerPassword = "vmtRemoteMediation"
	wsCommunicator.VmtServerAddress = VMTServerAddress
	wsCommunicator.LocalAddress = localAddress
	loginInfo = new(ConnectionInfo)
	loginInfo.Type = "Kubernetes"
	loginInfo.UserDefinedNameorIPAddress = "k8s_vmt"
	loginInfo.Username = "username"
	loginInfo.Password = "password"
	loginInfo.OpsManagerUsername = OpsManagerUsername
	loginInfo.OpsManagerPassword = OpsManagerPassword
	loginInfo.TargetIdentifier = TargetIdentifier
	msgHandler = new(MsgHandler)
	msgHandler.wscommunicator = wsCommunicator
	msgHandler.cInfo = loginInfo
	vMTApiRequestHandler = new(VMTApiRequestHandler)
	vMTApiRequestHandler.vmtServerAddr = wsCommunicator.VmtServerAddress
	vMTApiRequestHandler.opsManagerUsername = loginInfo.OpsManagerUsername
	vMTApiRequestHandler.opsManagerPassword = loginInfo.OpsManagerPassword
	msgHandler.vmtapi = vMTApiRequestHandler
	wsCommunicator.ServerMsgHandler = msgHandler

}

func main() {

	// Registration message is created and sent to VMTServer using network and authorization
	// parameters set in init() function
	containerInfo := msgHandler.CreateContainerInfo(wsCommunicator.LocalAddress)
	wsCommunicator.RegisterAndListen(containerInfo)

}
