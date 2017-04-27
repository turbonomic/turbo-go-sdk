package supplychain

import "github.com/turbonomic/turbo-go-sdk/pkg/proto"

const (
	getIPAddressMethodName string = "getAddress"

	serverIPProperty   = "UsesEndPoints"
	serverUUIDProperty = "Uuid"
)

func getIPHandler() *proto.ExternalEntityLink_PropertyHandler {
	directlyApply := false
	ipEntityType := proto.EntityDTO_IP
	methodName := getIPAddressMethodName

	return &proto.ExternalEntityLink_PropertyHandler{
		MethodName:    &methodName,
		DirectlyApply: &directlyApply,
		EntityType:    &ipEntityType,
	}
}

var VM_IP *proto.ExternalEntityLink_ServerEntityPropDef = getVirtualMachineIPProperty()

func getVirtualMachineIPProperty() *proto.ExternalEntityLink_ServerEntityPropDef {
	attribute := serverIPProperty
	vmEntityType := proto.EntityDTO_VIRTUAL_MACHINE

	return &proto.ExternalEntityLink_ServerEntityPropDef{
		Entity:          &vmEntityType,
		Attribute:       &attribute,
		PropertyHandler: getIPHandler(),
	}
}

var VM_UUID *proto.ExternalEntityLink_ServerEntityPropDef = getVirtualMachineUUIDProperty()

func getVirtualMachineUUIDProperty() *proto.ExternalEntityLink_ServerEntityPropDef {
	attribute := serverUUIDProperty
	vmEntityType := proto.EntityDTO_VIRTUAL_MACHINE

	return &proto.ExternalEntityLink_ServerEntityPropDef{
		Entity:    &vmEntityType,
		Attribute: &attribute,
	}
}
