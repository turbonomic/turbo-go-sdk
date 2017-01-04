package client

import (
	"github.com/turbonomic/turbo-go-sdk/example/sdk-example-probe/pkg/turboapi/api"
)

type Client struct {
	*RESTClient
}

func NewAPIClient(c *Config) *Client {
	restClient := NewRESTClient(c.ServerAddress, c.APIPath, c.BasicAuth)
	return &Client{restClient}
}

// Discover a target using api
// http://localhost:8400/vmturbo/api/targets/<name_or_address>
func (this *Client) DiscoverTarget(nameOrAddress string) error {
	_, err := this.Post().Resource(api.Resource_Type_Target).Name(nameOrAddress).Do()
	return err
}

// Add a ExampleProbe target to server
// example : http://localhost:8400/vmturbo/api/externaltargets?
//                     type=<target_type>&nameOrAddress=<host_address>&username=<username>&targetIdentifier=<target_identifier>&password=<password>
func (this *Client) AddTarget(targetType, nameOrAddress, targetIdentifier, username, password string) error {
	_, err := this.Post().Resource(api.Resource_Type_External_Target).
		Param("type", targetType).
		Param("nameOrAddress", nameOrAddress).
		Param("targetIdentifier", targetIdentifier).
		Param("username", username).
		Param("password", password).
		Do()

	return err
}
