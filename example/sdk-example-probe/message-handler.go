package main

import (
	"bytes"
	"github.com/golang/glog"

	"github.com/vmturbo/vmturbo-go-sdk/pkg/comm"
	"github.com/vmturbo/vmturbo-go-sdk/pkg/proto"
)

// This Struct is the implementation of communicator.ServerMessageHandler interface
type MsgHandler struct {
	wscommunicator *comm.WebSocketCommunicator
	cInfo          *ConnectionInfo
	vmtapi         *VMTApiRequestHandler
}

// Method used for adding a Target to a VMTServer
func (h *MsgHandler) AddTarget() {

	var requestDataB bytes.Buffer
	requestDataB.WriteString("?type=")
	requestDataB.WriteString(h.cInfo.Type)
	requestDataB.WriteString("&")
	requestDataB.WriteString("nameOrAddress=")
	requestDataB.WriteString(h.cInfo.UserDefinedNameorIPAddress)
	requestDataB.WriteString("&")
	requestDataB.WriteString("username=")
	requestDataB.WriteString(h.cInfo.Username)
	requestDataB.WriteString("&")
	requestDataB.WriteString("targetIdentifier=")
	requestDataB.WriteString(h.cInfo.TargetIdentifier)
	requestDataB.WriteString("&")
	requestDataB.WriteString("password=")
	requestDataB.WriteString(h.cInfo.Password)
	str := requestDataB.String()
	postReply, err := h.vmtapi.vmtApiPost("/externaltargets", str)
	if err != nil {
		glog.Infof(" postReply error")
	}

	if postReply.Status != "200 OK" {
		glog.Infof(" postReplyMessage error")
	}
}

// This Method validates our target which was previously added to the VMTServer
func (h *MsgHandler) Validate(serverMsg *proto.MediationServerMessage) {
	// messageID is a int32 , if nil then 0
	messageID := serverMsg.GetMessageID()
	validationResponse := new(proto.ValidationResponse)

	// creates a ClientMessageBuilder and sets ClientMessageBuilder.clientMessage.MessageID = messageID
	// sets clientMessage.ValidationResponse = validationResponse
	// type of clientMessage is MediationClientMessage
	clientMsg := comm.NewClientMessageBuilder(messageID).SetValidationResponse(validationResponse).Create()
	h.wscommunicator.SendClientMessage(clientMsg)
	glog.Infof("The client msg sent out is %++v", clientMsg)
	var requestDataB bytes.Buffer
	requestDataB.WriteString("?type=")
	requestDataB.WriteString(h.cInfo.Type)
	requestDataB.WriteString("&")
	requestDataB.WriteString("nameOrAddress=")
	requestDataB.WriteString(h.cInfo.UserDefinedNameorIPAddress)
	requestDataB.WriteString("&")
	requestDataB.WriteString("username=")
	requestDataB.WriteString(h.cInfo.Username)
	requestDataB.WriteString("&")
	requestDataB.WriteString("targetIdentifier=")
	requestDataB.WriteString(h.cInfo.TargetIdentifier)
	requestDataB.WriteString("&")
	requestDataB.WriteString("password=")
	requestDataB.WriteString(h.cInfo.Password)
	str := requestDataB.String()

	postReply, err := h.vmtapi.vmtApiPost("/targets", str)
	if err != nil {
		glog.Infof(" error in validate response from server")
		return
	}

	if postReply.Status != "200 OK" {
		glog.Infof("Validate reply came in with error")
	}
	return
}

func (h *MsgHandler) HandleAction(serverMsg *proto.MediationServerMessage) {
	glog.Infof("HandleAction called")
	return
}

// This Method sends all the topology entities and relationships found at
// this target to the VMTServer
func (h *MsgHandler) DiscoverTopology(serverMsg *proto.MediationServerMessage) {

	messageID := serverMsg.GetMessageID()
	simulatedProbe := &TargetProbe{}
	// add some fake nodes to simulatdProbe or just created it in getNodeEntityDTOs
	nodeEntityDTOs := simulatedProbe.getNodeEntityDTOs() // []*sdk.EntityDTO
	//  use simulated kubeclient to do ParseNode and ParsePod
	discoveryResponse := &proto.DiscoveryResponse{
		EntityDTO: nodeEntityDTOs,
	}
	clientMsg := comm.NewClientMessageBuilder(messageID).SetDiscoveryResponse(discoveryResponse).Create()
	h.wscommunicator.SendClientMessage(clientMsg)
	glog.Infof("The client msg sent out is %++v", clientMsg)
	return
}

// Function Creates ContainerInfo struct, sets Kubernetes Container Probe Information
// Returns pointer to newly created ContainerInfo
func (h *MsgHandler) CreateContainerInfo(localaddr string) *proto.ContainerInfo {
	var acctDefProps []*proto.AccountDefEntry
	targetIDAcctDefEntry := comm.NewAccountDefEntryBuilder(h.cInfo.TargetIdentifier,
		h.cInfo.UserDefinedNameorIPAddress, localaddr, ".*", proto.AccountDefEntry_OPTIONAL, false).Create()
	acctDefProps = append(acctDefProps, targetIDAcctDefEntry)
	usernameAcctDefEntry := comm.NewAccountDefEntryBuilder("username", "Username", h.cInfo.Username, ".*", proto.AccountDefEntry_OPTIONAL, false).Create()
	acctDefProps = append(acctDefProps, usernameAcctDefEntry)
	passwdAcctDefEntry := comm.NewAccountDefEntryBuilder("password", "Password", h.cInfo.Password, ".*", proto.AccountDefEntry_OPTIONAL, true).Create()
	acctDefProps = append(acctDefProps, passwdAcctDefEntry)
	//create the ProbeInfo struct with only type and category fields
	probeType := h.cInfo.Type
	probeCat := "Container"
	templateDTOs := createSupplyChain()
	probeInfo := comm.NewProbeInfoBuilder(probeType, probeCat, templateDTOs, acctDefProps).Create()
	// Create container
	containerInfo := new(proto.ContainerInfo)
	// Add probe to array of ProbeInfo* in container
	probes := append(containerInfo.Probes, probeInfo)
	containerInfo.Probes = probes
	return containerInfo
}
