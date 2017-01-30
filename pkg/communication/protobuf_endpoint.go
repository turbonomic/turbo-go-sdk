package communication

import (
	"fmt"

	goproto "github.com/golang/protobuf/proto"
	"github.com/golang/glog"

	proto "github.com/turbonomic/turbo-go-sdk/pkg/proto"
	version "github.com/turbonomic/turbo-go-sdk/pkg/version"
)

// Endpoint to handle communication of a particular protobuf message type with the server
type ProtobufEndpoint interface {
	GetName() string
	CloseEndpoint()
	Send(messageToSend *EndpointMessage)
	GetMessageHandler() ProtobufMessage
	MessageReceiver() chan goproto.Message
	//AddEventHandler(handler EndpointEventHandler)
}


type EndpointMessage struct {
	ProtobufMessage goproto.Message
}

// =====================================================================================
// Parser interface for different server messages
type ProtobufMessage interface {
	parse(rawMsg []byte)
	GetMessage() goproto.Message
}

// Parser for all the Mediation Requests such as Discovery, Validation, Action etc
type ServerRequest struct {
	ServerMsg *proto.MediationServerMessage
}

// Parser for the Negotiation Response
type NegotiationResponse struct {
	NegotiationMsg *version.NegotiationAnswer
}

// Parser for the Registration Response
type RegistrationResponse struct {
	RegistrationMsg *proto.Ack
}

func (sr *ServerRequest) GetMessage() goproto.Message {
	return sr.ServerMsg
}

func (sr *ServerRequest) parse(rawMsg []byte) {
	fmt.Printf("[ServerRequest] parsing %s", rawMsg)
	fmt.Println("")
	// Parse the input stream
	serverMsg := &proto.MediationServerMessage{}
	err := goproto.Unmarshal(rawMsg, serverMsg)
	if err != nil {
		glog.Error("[ServerRequest] Error unmarshalling transprot input stream to protobuf message, please make sure you are running the latest VMT server")
		glog.Fatal("[ServerRequest] unmarshaling error: ", err)
	}
	sr.ServerMsg = serverMsg
	return
}

func (nr *NegotiationResponse) GetMessage() goproto.Message {
	return nr.NegotiationMsg
}

func (nr *NegotiationResponse) parse(rawMsg []byte) {
	fmt.Printf("[NegotiationResponse] parsing %s", rawMsg)
	fmt.Println("")
	// Parse the input stream
	serverMsg := &version.NegotiationAnswer{}
	err := goproto.Unmarshal(rawMsg, serverMsg)
	if err != nil {
		glog.Error("[NegotiationResponse] Error unmarshalling transprot input stream to protobuf message, please make sure you are running the latest VMT server")
		glog.Fatal("[NegotiationResponse] unmarshaling error: ", err)
	}
	nr.NegotiationMsg = serverMsg
	return
}

func (rr *RegistrationResponse) GetMessage() goproto.Message {
	return rr.RegistrationMsg
}

func (rr *RegistrationResponse) parse(rawMsg []byte) {
	fmt.Printf("[RegistrationResponse] parsing %s", rawMsg)
	fmt.Println("")
	// Parse the input stream
	serverMsg := &proto.Ack{}
	err := goproto.Unmarshal(rawMsg, serverMsg)
	if err != nil {
		glog.Error("[RegistrationResponse] Error unmarshalling transprot input stream to protobuf message, please make sure you are running the latest VMT server")
		glog.Fatal("[RegistrationResponse] unmarshaling error: ", err)
	}
	rr.RegistrationMsg = serverMsg
	return
}
