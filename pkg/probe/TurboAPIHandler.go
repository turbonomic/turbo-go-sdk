package probe

import (
	"bytes"
	"fmt"

	"github.com/golang/glog"

	api "github.com/turbonomic/turbo-go-sdk/pkg/vmtapi"
)
type TurboAPIConfig struct {
	VmtServerAddress	string
	VmtUser			string
	VmtPassword 		string
}

type TurboAPIHandler struct {

	TurboAPIClient	*api.VmtApi
 	// map of specific handlers
}

func NewTurboAPIHandler(conf *TurboAPIConfig) *TurboAPIHandler {
	handler := &TurboAPIHandler{}

	apiClient := api.NewVmtApi(conf.VmtServerAddress, conf.VmtUser, conf.VmtPassword)
	handler.TurboAPIClient = apiClient
	return handler
}

// Use the vmt restAPI to add a Turbo target.
func (handler *TurboAPIHandler) AddTarget(target *TurboTarget) error {
	// TODO: Check if the Target already exists in the server ?
	targetType := target.GetTargetType()
	targetIdentifier := target.GetTargetId()
	//nameOrAddress := target.GetNameOrAddress()
	username := target.GetUser()
	password := target.GetPassword()

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

	requestData["username"] = username
	requestDataBuffer.WriteString("username=")
	requestDataBuffer.WriteString(username)
	requestDataBuffer.WriteString("&")

	requestData["targetIdentifier"] = targetIdentifier
	requestDataBuffer.WriteString("targetIdentifier=")
	requestDataBuffer.WriteString(targetIdentifier)
	requestDataBuffer.WriteString("&")

	requestData["password"] = password
	requestDataBuffer.WriteString("password=")
	requestDataBuffer.WriteString(password)

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
func (handler *TurboAPIHandler) DiscoverTarget(target *TurboTarget) error {

	// Discover Mesos target.
	glog.V(3).Info("Calling VMTurbo REST API to initiate a new discovery.")

	respMsg, err := handler.TurboAPIClient.Post("/targets/"+target.GetNameOrAddress(), "")
	if err != nil {
		return err
	}
	fmt.Println("Discover target response is %s", respMsg)

	return nil
}
