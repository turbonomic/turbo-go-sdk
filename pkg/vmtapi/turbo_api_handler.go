package vmtapi

import "fmt"

import (
	"bytes"

	"github.com/golang/glog"

	"github.com/turbonomic/turbo-go-sdk/pkg/probe"
	"github.com/vmturbo/vmturbo-go-sdk/pkg/vmtapi"
)


type TurboAPIConfig struct {
	VmtRestServerAddress string
	VmtRestUser          string
	VmtRestPassword      string
}

// TODO:
func (apiConfig *TurboAPIConfig) ValidateTurboAPIConfig() bool {
	fmt.Println("========== Turbo Rest API Config =============")
	fmt.Println("VmtServerAddress : " + string(apiConfig.VmtRestServerAddress))
	fmt.Println("VmtUsername : " + apiConfig.VmtRestUser)
	fmt.Println("VmtPassword : " + apiConfig.VmtRestPassword)
	return true
}

// =====================================================================================================================

type TurboAPIHandler struct {

	TurboAPIClient	*vmtapi.VmtApi
 	// map of specific handlers
}

func NewTurboAPIHandler(conf *TurboAPIConfig) *TurboAPIHandler {
	fmt.Println("---------- Created TurboAPIHandler ----------")
	handler := &TurboAPIHandler{}

	apiClient := vmtapi.NewVmtApi(conf.VmtRestServerAddress, conf.VmtRestUser, conf.VmtRestPassword)
	handler.TurboAPIClient = apiClient
	return handler
}

// Use the vmt restAPI to add a Turbo target.
func (handler *TurboAPIHandler) AddTurboTarget(target *probe.TurboTarget) error {
	// TODO: Check if the Target already exists in the server ?
	targetType := target.GetTargetType()
	//targetIdentifier := target.GetTargetId()
	//nameOrAddress := target.GetNameOrAddress()
	//username := target.GetUser()
	//password := target.GetPassword()

	fmt.Println("[TurboAPIHandler] Calling VMTurbo REST API to added current %s target.", targetType)
	// Create request string parameters
	requestData := make(map[string]string)

	var requestDataBuffer bytes.Buffer

	requestData["type"] = targetType
	requestDataBuffer.WriteString("?type=")
	requestDataBuffer.WriteString(targetType)
	requestDataBuffer.WriteString("&")

	//requestData["nameOrAddress"] = nameOrAddress
	//requestDataBuffer.WriteString("nameOrAddress=")
	//requestDataBuffer.WriteString(nameOrAddress)
	//requestDataBuffer.WriteString("&")

	//requestData["username"] = username
	//requestDataBuffer.WriteString("username=")
	//requestDataBuffer.WriteString(username)
	//requestDataBuffer.WriteString("&")
	//
	//requestData["targetIdentifier"] = targetIdentifier
	//requestDataBuffer.WriteString("targetIdentifier=")
	//requestDataBuffer.WriteString(targetIdentifier)
	//requestDataBuffer.WriteString("&")
	//
	//requestData["password"] = password
	//requestDataBuffer.WriteString("password=")
	//requestDataBuffer.WriteString(password)
	//requestDataBuffer.WriteString("&")

	acctVals := target.AccountValues
	for idx, acctEntry := range acctVals {
		prop := *acctEntry.Key
		requestData[prop] = prop
		requestDataBuffer.WriteString(prop+"=")
		requestDataBuffer.WriteString(*acctEntry.StringValue)

		if idx != (len(acctVals)-1) {
			requestDataBuffer.WriteString("&")
		}
	}


	s := requestDataBuffer.String()

	// Create HTTP Endpoint to send and handle target addition messages
	respMsg, err := handler.TurboAPIClient.Post("/externaltargets", s)
	if err != nil {
		fmt.Println(" ERROR: %s", err)
		return err
	}
	fmt.Println("[TurboAPIHandler] Add target response is %s", respMsg)
	return nil
}

// Send an API request to make server start a discovery process on current Mesos
func (handler *TurboAPIHandler) DiscoverTarget(target *probe.TurboTarget) error {

	// Discover Mesos target.
	glog.V(3).Info("Calling VMTurbo REST API to initiate a new discovery.")

	respMsg, err := handler.TurboAPIClient.Post("/targets/"+target.GetNameOrAddress(), "")
	if err != nil {
		return err
	}
	fmt.Println("Discover target response is %s", respMsg)

	return nil
}
