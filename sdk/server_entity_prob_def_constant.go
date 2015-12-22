package sdk

var ipHandler *ExternalEntityLink_PropertyHandler = getIpHandler()

func getIpHandler() *ExternalEntityLink_PropertyHandler {
	directlyApply := false
	ipEntityType := EntityDTO_IP
	METHOD_NAME_GET_IP_ADDRESS := "getAddress"

	return &ExternalEntityLink_PropertyHandler{
		MethodName:    &METHOD_NAME_GET_IP_ADDRESS,
		DirectlyApply: &directlyApply,
		EntityType:    &ipEntityType,
	}
}

var VM_IP *ExternalEntityLink_ServerEntityPropDef = getVirtualMachineIpProperty()

func getVirtualMachineIpProperty() *ExternalEntityLink_ServerEntityPropDef {
	attribute := "UsesEndPoints"
	vmEntityType := EntityDTO_VIRTUAL_MACHINE

	return &ExternalEntityLink_ServerEntityPropDef{
		Entity:          &vmEntityType,
		Attribute:       &attribute,
		PropertyHandler: ipHandler,
	}
}
