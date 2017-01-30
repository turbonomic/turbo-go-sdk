package communication

import (
	"fmt"

	"github.com/golang/glog"

	proto "github.com/turbonomic/turbo-go-sdk/pkg/proto"
	version "github.com/turbonomic/turbo-go-sdk/pkg/version"
)

type SdkClientProtocol struct {
	allProbes map[string]*ProbeProperties
	TransportReady chan bool
}

func CreateSdkClientProtocolHandler(allProbes map[string]*ProbeProperties, done chan bool) (*SdkClientProtocol) {
	return &SdkClientProtocol{
		allProbes: allProbes,
		TransportReady: done,	//(chan string),
	}
}

func (clientProtocol *SdkClientProtocol) handleClientProtocol(transport ITransport) {
	fmt.Println("[SdkClientProtocol] Starting Protocol Negotiation ....")
	status := clientProtocol.NegotiateVersion(transport)

	if !status {
		fmt.Println("[SdkClientProtocol] Failure during Protocol Negotiation, Registration message will not be sent")
		clientProtocol.TransportReady <- false
		return
	}
	fmt.Println("[SdkClientProtocol] Starting Probe Registration ....")
	status = clientProtocol.HandleRegistration(transport)
	if !status {
		fmt.Println("[SdkClientProtocol] Failure during Registration, cannot receive server messages")
		clientProtocol.TransportReady <- false
		return
	}

	clientProtocol.TransportReady <- true
}

// TODO: Need something back from the protocol to say that the client is ready

// ============================== Protocol Version Negotiation =========================
func (clientProtocol *SdkClientProtocol) NegotiateVersion(transport ITransport) bool {
	versionStr := getSpecificationVersion()
	request := &version.NegotiationRequest {
			ProtocolVersion: &versionStr,
		}
	glog.V(3).Infof("Send negotiation message: %+v", request)

	// Create Protobuf Endpoint to send and handle negotiation messages
	protoMsg := &NegotiationResponse {}	// handler for the response
	endpoint := CreateClientProtobufEndpoint2("NegotiationEndpoint", transport, protoMsg)

	endMsg := &EndpointMessage{
		ProtobufMessage: request,
	}
	endpoint.Send(endMsg)
	defer close(endpoint.MessageReceiver())
	// Wait for the response to be received by the transport and then parsed and put on the endpoint's message channel
	serverMsg, ok := <-endpoint.MessageReceiver()
	if !ok {
		fmt.Println("[SdkClientProtocol][" + endpoint.GetName() + "] : Endpoint Receiver channel is closed")
		endpoint.CloseEndpoint()
		return false
	}
	fmt.Printf("[SdkClientProtocol] [" + endpoint.GetName() + "] : Received: %s\n", serverMsg)

	// Handler response
	negotiationResponse := protoMsg.NegotiationMsg
	negotiationResponse.GetNegotiationResult()

	if negotiationResponse.GetNegotiationResult().String() != version.NegotiationAnswer_ACCEPTED.String() {
		fmt.Println("[SdkClientProtocol] Protocol version negotiation failed",
				negotiationResponse.GetNegotiationResult().String() + ") :" + negotiationResponse.GetDescription())
		return false
	}
	fmt.Println("[SdkClientProtocol] Protocol version is accepted by server: {}", negotiationResponse.GetDescription())
	endpoint.CloseEndpoint()
	return true
}

func getSpecificationVersion() (string) {
	specificationVersion := "5.8.0-SNAPSHOT"
	//TODO: final String specificationVersion = CommonDTO.class.getPackage().getSpecificationVersion();
	if specificationVersion == "" {
		glog.V(3).Infof("Specification-Version is null for CommonDTO class. ",
				 "This should not be so in production run")
		return "";
	}
	return specificationVersion
}


//type NegotiationEndpointEventHandler struct {
//	MessageChannel chan goproto.Message
//}

//func (eh *NegotiationEndpointEventHandler) onClose() {}
//func (eh *NegotiationEndpointEventHandler) onMessage(handler *EndpointEventHandler) {}

// ======================= Registration ============================
// Send registration message
func (clientProtocol *SdkClientProtocol) HandleRegistration(transport ITransport) bool {
	containerInfo, err := clientProtocol.MakeContainerInfo()
	if  err != nil {
		glog.Error("Error creating ContainerInfo")
		return false
	}

	glog.V(3).Infof("[SdkClientProtocol] Send registration message: %+v", containerInfo)

	// Create Protobuf Endpoint to send and handle registration messages
	protoMsg := &RegistrationResponse{}
	endpoint := CreateClientProtobufEndpoint2("RegistrationEndpoint", transport, protoMsg)

	endMsg := &EndpointMessage{
		ProtobufMessage: containerInfo,
	}
	endpoint.Send(endMsg)
	fmt.Println("[SdkClientProtocol][" + endpoint.GetName() + "] : Waiting for registration response")
	// Wait for the response to be received by the transport and then parsed and put on the endpoint's message channel
	serverMsg, ok := <-endpoint.MessageReceiver()
	if !ok {
		fmt.Println("[" + endpoint.GetName() + "] : Endpoint Receiver channel is closed")
		endpoint.CloseEndpoint()
		return false
	}
	fmt.Printf("[SdkClientProtocol] [" + endpoint.GetName() + "] : Received: %s\n", serverMsg)

	// Handler response
	registrationResponse := protoMsg.RegistrationMsg
	if registrationResponse == nil {
		fmt.Println("[SdkClientProtocol] Probe registration failed, null ack")
		return false
	}
	endpoint.CloseEndpoint()

	return true
}

func (clientProtocol *SdkClientProtocol) MakeContainerInfo() (*proto.ContainerInfo, error) {
	// 4. Add example probe to probeInfo list, here it is the only probe supported.
	var probes []*proto.ProbeInfo

	for k,v := range clientProtocol.allProbes {
		fmt.Println("SdkClientProtocol] Creating Probe Info for", k)
		turboProbe := v.Probe
		var probeInfo *proto.ProbeInfo
		var err error
		probeInfo, err = turboProbe.GetProbeInfo()

		if err != nil {
			return nil, err
		}
		probes = append(probes, probeInfo)
	}

	return &proto.ContainerInfo{
		Probes: probes,
	}, nil
}


// TODO: remove this from the test probes that use the IProbe interface
//func buildProbeInfo(probeProps *ProbeProperties) (*proto.ProbeInfo, error) {
//	// 1. Get the account definition for probe
//	var acctDefProps []*proto.AccountDefEntry
//	var templateDtos  []*proto.TemplateDTO
//
//	turboProbe := probeProps.Probe
//
//	acctDefProps = turboProbe.RegistrationClient.GetAccountDefinition()
//
//	// 2. Get the supply chain.
//	//templateDtos := probeInterface.GetSupplyChainDefinition()	//&supplychain.SupplyChainFactory{}
//	templateDtos = turboProbe.RegistrationClient.GetSupplyChainDefinition()
//
//	// 3. construct the example probe info.
//	probeConfig := probeProps.ProbeSignature
//	probeCat := probeConfig.ProbeCategory	//"Container"
//	probeType := probeConfig.ProbeType
//	probeInfo := builder.NewProbeInfoBuilder(probeType, probeCat, templateDtos, acctDefProps).Create()
//	id := "targetIdentifier"		// TODO: parameterize this for different probes
//	probeInfo.TargetIdentifierField = &id
//
//	// 4. Add example probe to probeInfo list, here it is the only probe supported.
//	//var probes []*proto.ProbeInfo
//	//probes = append(probes, exampleProbe)
//
//	return probeInfo, nil
//}

