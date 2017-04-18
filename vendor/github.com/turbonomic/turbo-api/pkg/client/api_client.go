package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"crypto/tls"
	"errors"

	"github.com/turbonomic/turbo-api/pkg/api"
)

type Client struct {
	*RESTClient
}

// Create a Turbo API Client based on basic authentication.
func NewAPIClientWithBA(c *Config) (*Client, error) {
	if c.basicAuth == nil {
		return nil, errors.New("Basic authentication is not set")
	}
	client := http.DefaultClient
	// If use https, disable the security check.
	if c.serverAddress.Scheme == "https" {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}
	restClient := NewRESTClient(client, c.serverAddress, c.apiPath).BasicAuthentication(c.basicAuth)
	return &Client{restClient}, nil
}

// Discover a target using API
func (c *Client) DiscoverTarget(uuid string) (*Result, error) {
	response, err := c.Post().Resource(api.Resource_Type_Target).Name(uuid).Do()
	if err != nil {
		return nil, fmt.Errorf("Failed to discover target %s: %s", uuid, err)
	}
	if response.statusCode != 200 {
		return nil, fmt.Errorf("Unsuccessful discovery response: %s", response.status)
	}
	return &response, nil
}

//Add a target usign API
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
		return nil, fmt.Errorf("Unsuccessful target addition response: %s", response.status)
	}
	return &response, nil
}
