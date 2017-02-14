package vmtapi

import "fmt"

import (
	"bytes"

	"github.com/golang/glog"
	"github.com/turbonomic/turbo-go-sdk/pkg/probe"
	"errors"
)

type TurboAPIConfig struct {
	VmtRestServerAddress string
	VmtRestUser          string
	VmtRestPassword      string
}

func (apiConfig *TurboAPIConfig) ValidateTurboAPIConfig() (bool, error) {
	if apiConfig.VmtRestServerAddress == "" {
		return false, errors.New("Turbo server ip address is required "+ fmt.Sprint(apiConfig))
	}
	if apiConfig.VmtRestUser == "" {
		return false, errors.New("Turbo Rest API user is required "+ fmt.Sprint(apiConfig))
	}

	if apiConfig.VmtRestUser == "" {
		return false, errors.New("Turbo Rest API user password is required "+ fmt.Sprint(apiConfig))
	}


	fmt.Println("========== Turbo Rest API Config =============")
	fmt.Println("VmtServerAddress : " + string(apiConfig.VmtRestServerAddress))
	fmt.Println("VmtUsername : " + apiConfig.VmtRestUser)
	fmt.Println("VmtPassword : " + apiConfig.VmtRestPassword)
	// TODO:
	return true, nil
}

// =====================================================================================================================

type TurboAPIHandler struct {
	TurboAPIClient *VmtApi
	// map of specific handlers
}

func NewTurboAPIHandler(conf *TurboAPIConfig) *TurboAPIHandler {
	glog.Infof("---------- Created TurboAPIHandler ----------")
	handler := &TurboAPIHandler{}

	apiClient := NewVmtApi(conf.VmtRestServerAddress, conf.VmtRestUser, conf.VmtRestPassword)
	handler.TurboAPIClient = apiClient
	return handler
}

// Use the vmt restAPI to add a Turbo target.
func (handler *TurboAPIHandler) AddTurboTarget(target *probe.TurboTarget) error {
	// TODO: Check if the Target already exists in the server ?
	targetType := target.GetTargetType()
	glog.Infof("Calling VMTurbo REST API to added current %s target.", targetType)

	// Create request string parameters
	requestData := make(map[string]string)

	var requestDataBuffer bytes.Buffer

	requestData["type"] = targetType
	requestDataBuffer.WriteString("?type=")
	requestDataBuffer.WriteString(targetType)
	requestDataBuffer.WriteString("&")

	acctVals := target.AccountValues
	for idx, acctEntry := range acctVals {
		prop := *acctEntry.Key
		requestData[prop] = prop
		requestDataBuffer.WriteString(prop + "=")
		requestDataBuffer.WriteString(*acctEntry.StringValue)

		if idx != (len(acctVals) - 1) {
			requestDataBuffer.WriteString("&")
		}
	}

	s := requestDataBuffer.String()

	// Create HTTP Endpoint to send and handle target addition messages
	respMsg, err := handler.TurboAPIClient.Post("/externaltargets", s)
	if err != nil {
		glog.Errorf(" ERROR: %s", err)
		return err
	}
	glog.Infof("Add target response is %s", respMsg)
	return nil
}

// Send an API request to make server start a discovery process on current Mesos
func (handler *TurboAPIHandler) DiscoverTarget(target *probe.TurboTarget) error {

	// Discover Turbo target.
	glog.V(3).Info("Calling VMTurbo REST API to initiate a new discovery.")

	respMsg, err := handler.TurboAPIClient.Post("/targets/" + target.GetNameOrAddress(), "")
	if err != nil {
		return err
	}
	glog.Infof("Discover target response is %s", respMsg)

	return nil
}
