package client

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/turbonomic/turbo-api/pkg/api"
)

type Client struct {
	*RESTClient
}

// Create a Turbo API Client based on basic authentication.
func NewAPIClientWithBA(c *Config) (*Client, error) {
	if c.BasicAuth == nil {
		return nil, fmt.Errorf("Basic authentication is not set")
	}
	restClient := NewRESTClient(http.DefaultClient, c.ServerAddress, c.APIPath).BasicAuthentication(c.BasicAuth)
	return &Client{restClient}, nil
}

// Discover a target using api
// <turbo_server_address>/vmturbo/api/targets/<uuid>
func (c *Client) DiscoverTarget(uuid string) (Result, error) {
	response, err := c.Post().Resource(api.Resource_Type_Target).Name(uuid).Do()
	if err != nil {
		return Result{}, fmt.Errorf("Failed to discover target %s: %s", uuid, err)
	}
	return response, nil
}

func (c *Client) AddTarget(target *api.Target) (Result, error) {
	targetData, err := json.Marshal(target)
	if err != nil {
		return Result{}, err
	}
	response, err := c.Post().Resource(api.Resource_Type_Target).
		Header("Content-Type", "application/json").
		Header("Accept", "application/json").
		Data(targetData).
		Do()
	if err != nil {
		return Result{}, err
	}
	return response, nil
}

// Add a ExampleProbe target to server
// example : <turbo_server_address>/vmturbo/api/externaltargets?
//                     type=<target_type>&nameOrAddress=<host_address>&username=<username>&targetIdentifier=<target_identifier>&password=<password>
func (c *Client) AddExternalTarget(exTarget *api.ExternalTarget) (Result, error) {
	res, err := c.Post().Resource(api.Resource_Type_External_Target).
		Param("type", exTarget.TargetType).
		Param("nameOrAddress", exTarget.NameOrAddress).
		Param("targetIdentifier", exTarget.TargetIdentifier).
		Param("username", exTarget.Username).
		Param("password", exTarget.Password).
		Do()
	if err != nil {
		return Result{}, fmt.Errorf("Failed to add external target: %s", err)
	}
	fmt.Printf("Response is %s\n", res)
	return res, nil
}
