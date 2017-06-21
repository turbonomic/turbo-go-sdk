package version

// Remote Mediation Clients are expected to send the protobuf message version to ensure compatibility with the server.
// This version string is sent as part of the registration protocol in the Negotiation message. Only if the version is
// accepted by the server then the client probes are allowed to register with the server
const (
	PROTOBUF_VERSION string = "5.9.0"
)
