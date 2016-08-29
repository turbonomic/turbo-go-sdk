package common

import "github.com/vmturbo/vmturbo-go-sdk/pkg/proto"

var ipHandler *proto.ExternalEntityLink_PropertyHandler = getIpHandler()

func getIpHandler() *proto.ExternalEntityLink_PropertyHandler {
	directlyApply := false
	ipEntityType := proto.EntityDTO_IP
	METHOD_NAME_GET_IP_ADDRESS := "getAddress"

	return &proto.ExternalEntityLink_PropertyHandler{
		MethodName:    &METHOD_NAME_GET_IP_ADDRESS,
		DirectlyApply: &directlyApply,
		EntityType:    &ipEntityType,
	}
}

var VM_IP *proto.ExternalEntityLink_ServerEntityPropDef = getVirtualMachineIpProperty()

func getVirtualMachineIpProperty() *proto.ExternalEntityLink_ServerEntityPropDef {
	attribute := "UsesEndPoints"
	vmEntityType := proto.EntityDTO_VIRTUAL_MACHINE

	return &proto.ExternalEntityLink_ServerEntityPropDef{
		Entity:          &vmEntityType,
		Attribute:       &attribute,
		PropertyHandler: ipHandler,
	}
}
