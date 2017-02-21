package mediationcontainer

import (

	"errors"
	"github.com/golang/glog"
	goproto "github.com/golang/protobuf/proto"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/pkg/version"
)

// Endpoint to handle communication of a particular protobuf message type with the server
type ProtobufEndpoint interface {
	GetName() string
	GetTransport() ITransport
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
	parse(rawMsg []byte) error
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

func (sr *ServerRequest) parse(rawMsg []byte) (error) {
	glog.V(2).Infof("Parsing %s\n", rawMsg)
	// Parse the input stream
	serverMsg := &proto.MediationServerMessage{}
	err := goproto.Unmarshal(rawMsg, serverMsg)
	if err != nil {
		glog.Error("[ServerRequest] Error unmarshalling transport input stream to protobuf message, please make sure you are running the latest VMT server")
		glog.Error("[ServerRequest] unmarshaling error: ", err)
		return errors.New("[ServerRequest] Error unmarshalling transport input stream to protobuf message : " + err.Error())
	}
	sr.ServerMsg = serverMsg
	return nil
}

func (nr *NegotiationResponse) GetMessage() goproto.Message {
	return nr.NegotiationMsg
}

func (nr *NegotiationResponse) parse(rawMsg []byte) error {
	glog.V(2).Infof("Parsing %s\n", rawMsg)
	// Parse the input stream
	serverMsg := &version.NegotiationAnswer{}
	err := goproto.Unmarshal(rawMsg, serverMsg)
	if err != nil {
		glog.Error("[NegotiationResponse] Error unmarshalling transport input stream to protobuf message, please make sure you are running the latest VMT server")
		glog.Error("[NegotiationResponse] unmarshaling error: ", err)
		return errors.New("[NegotiationResponse] Error unmarshalling transport input stream to protobuf message : "  + err.Error())
	}
	nr.NegotiationMsg = serverMsg
	return nil
}

func (rr *RegistrationResponse) GetMessage() goproto.Message {
	return rr.RegistrationMsg
}

func (rr *RegistrationResponse) parse(rawMsg []byte) error {
	glog.V(2).Infof("Parsing %s\n", rawMsg)
	// Parse the input stream
	serverMsg := &proto.Ack{}
	err := goproto.Unmarshal(rawMsg, serverMsg)
	if err != nil {
		glog.Error("[RegistrationResponse] Error unmarshalling transport input stream to protobuf message, please make sure you are running the latest VMT server")
		glog.Error("[RegistrationResponse] unmarshaling error: ", err)
		return errors.New("[RegistrationResponse] Error unmarshalling transport input stream to protobuf message : " +  err.Error())
	}
	rr.RegistrationMsg = serverMsg
	return nil
}
