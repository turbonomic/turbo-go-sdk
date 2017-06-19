package mediationcontainer

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/pkg/version"

	"github.com/golang/glog"
)

type SdkClientProtocol struct {
	allProbes map[string]*ProbeProperties
	version   string
	//TransportReady chan bool
}

func CreateSdkClientProtocolHandler(allProbes map[string]*ProbeProperties, version string) *SdkClientProtocol {
	return &SdkClientProtocol{
		allProbes: allProbes,
		version:   version,
		//TransportReady: done,
	}
}

func (clientProtocol *SdkClientProtocol) handleClientProtocol(transport ITransport, transportReady chan bool) {
	glog.V(2).Infof("Starting Protocol Negotiation ....")
	status := clientProtocol.NegotiateVersion(transport)

	if !status {
		glog.Errorf("Failure during Protocol Negotiation, Registration message will not be sent")
		transportReady <- false
		// clientProtocol.TransportReady <- false
		return
	}
	glog.V(2).Infof("[SdkClientProtocol] Starting Probe Registration ....")
	status = clientProtocol.HandleRegistration(transport)
	if !status {
		glog.Errorf("Failure during Registration, cannot receive server messages")
		transportReady <- false
		// clientProtocol.TransportReady <- false
		return
	}

	transportReady <- true
	// clientProtocol.TransportReady <- true
}

// ============================== Protocol Version Negotiation =========================

func (clientProtocol *SdkClientProtocol) NegotiateVersion(transport ITransport) bool {
	versionStr := clientProtocol.version
	request := &version.NegotiationRequest{
		ProtocolVersion: &versionStr,
	}
	glog.V(3).Infof("Send negotiation message: %+v", request)

	// Create Protobuf Endpoint to send and handle negotiation messages
	protoMsg := &NegotiationResponse{} // handler for the response
	endpoint := CreateClientProtoBufEndpoint("NegotiationEndpoint", transport, protoMsg, true)

	endMsg := &EndpointMessage{
		ProtobufMessage: request,
	}
	endpoint.Send(endMsg)
	//defer close(endpoint.MessageReceiver())
	// Wait for the response to be received by the transport and then parsed and put on the endpoint's message channel
	serverMsg, ok := <-endpoint.MessageReceiver()
	if !ok {
		glog.Errorf("[%s] : Endpoint Receiver channel is closed", endpoint.GetName())
		endpoint.CloseEndpoint()
		return false
	}
	glog.V(3).Infof("["+endpoint.GetName()+"] : Received: %s\n", serverMsg)

	// Handler response
	negotiationResponse := protoMsg.NegotiationMsg
	if negotiationResponse == nil {
		glog.Error("Probe Protocol failed, null negotiation response")
		endpoint.CloseEndpoint()
		return false
	}
	negotiationResponse.GetNegotiationResult()

	if negotiationResponse.GetNegotiationResult().String() != version.NegotiationAnswer_ACCEPTED.String() {
		glog.Errorf("Protocol version negotiation failed",
			negotiationResponse.GetNegotiationResult().String()+") :"+negotiationResponse.GetDescription())
		return false
	}
	glog.V(4).Infof("[SdkClientProtocol] Protocol version is accepted by server: %s", negotiationResponse.GetDescription())
	endpoint.CloseEndpoint()
	return true
}

// ======================= Registration ============================
// Send registration message
func (clientProtocol *SdkClientProtocol) HandleRegistration(transport ITransport) bool {
	containerInfo, err := clientProtocol.MakeContainerInfo()
	if err != nil {
		glog.Error("Error creating ContainerInfo")
		return false
	}

	glog.V(3).Infof("Send registration message: %+v", containerInfo)

	// Create Protobuf Endpoint to send and handle registration messages
	protoMsg := &RegistrationResponse{}
	endpoint := CreateClientProtoBufEndpoint("RegistrationEndpoint", transport, protoMsg, true)

	endMsg := &EndpointMessage{
		ProtobufMessage: containerInfo,
	}
	endpoint.Send(endMsg)
	//defer close(endpoint.MessageReceiver())
	// Wait for the response to be received by the transport and then parsed and put on the endpoint's message channel
	serverMsg, ok := <-endpoint.MessageReceiver()
	if !ok {
		glog.Errorf("[HandleRegistration] [" + endpoint.GetName() + "] : Endpoint Receiver channel is closed")
		endpoint.CloseEndpoint()
		return false
	}
	glog.V(2).Infof("[HandleRegistration] ["+endpoint.GetName()+"] : Received: %s\n", serverMsg)

	// Handler response
	registrationResponse := protoMsg.RegistrationMsg
	if registrationResponse == nil {
		glog.Errorf("Probe registration failed, null ack")
		return false
	}
	endpoint.CloseEndpoint()

	return true
}

func (clientProtocol *SdkClientProtocol) MakeContainerInfo() (*proto.ContainerInfo, error) {
	var probes []*proto.ProbeInfo

	for k, v := range clientProtocol.allProbes {
		glog.V(2).Infof("SdkClientProtocol] Creating Probe Info for", k)
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
