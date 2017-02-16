package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"crypto/tls"
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
	client := http.DefaultClient
	// If use https, disable the security check.
	if c.ServerAddress.Scheme == "https" {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}
	restClient := NewRESTClient(client, c.ServerAddress, c.APIPath).BasicAuthentication(c.BasicAuth)
	return &Client{restClient}, nil
}

// Discover a target using API
// <turbo_server_address>/vmturbo/api/targets/<uuid>
func (c *Client) DiscoverTarget(uuid string) (*Result, error) {
	response, err := c.Post().Resource(api.Resource_Type_Target).Name(uuid).Do()
	if err != nil {
		return nil, fmt.Errorf("Failed to discover target %s: %s", uuid, err)
	}
	if response.statusCode != 200 {
		return nil, fmt.Errorf("Unsuccessful discovery response: %v", response)
	}
	return &response, nil
}

// Add a target using API.
// <turbo_server_address>/vmturbo/rest/targets
func (c *Client) AddTarget(target *api.Target) (*Result, error) {
	targetData, err := json.Marshal(target)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshall target instance: %s", err)
	}
	response, err := c.Post().Resource(api.Resource_Type_Target).
		Header("Content-Type", "application/json").
		Header("Accept", "application/json").
		Data(targetData).
		Do()
	if err != nil {
		return nil, fmt.Errorf("Failed to add target: %s", err)
	}
	if response.statusCode != 200 {
		return nil, fmt.Errorf("Unsuccessful target addition response: %v", response)
	}
	return &response, nil
}
