package service

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/mediationcontainer"
	"github.com/turbonomic/turbo-go-sdk/util/rand"
	"testing"
)

func TestValidateRestAPIConfig(t *testing.T) {
	table := []struct {
		user     string
		password string

		expectErr bool
	}{
		{
			user:      "",
			password:  rand.String(5),
			expectErr: true,
		},
		{
			user:      rand.String(10),
			password:  rand.String(10),
			expectErr: false,
		},
		{
			user:      rand.String(10),
			password:  "",
			expectErr: true,
		},
	}
	for _, item := range table {
		config := &RestAPIConfig{
			OpsManagerUsername: item.user,
			OpsManagerPassword: item.password,
		}
		err := config.ValidRestAPIConfig()
		if item.expectErr {
			if err == nil {
				t.Errorf("Expected error, got no error. %s", config)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		}
	}
}

func TestValidateTurboCommunicationConfig(t *testing.T) {
	table := []struct {
		config    *TurboCommunicationConfig
		expectErr bool
	}{
		{
			config: &TurboCommunicationConfig{
				mediationcontainer.ServerMeta{
					TurboServer: "invalid",
				},
				mediationcontainer.WebSocketConfig{
					LocalAddress:      "http://127.0.0.1",
					WebSocketUsername: rand.String(10),
					WebSocketPassword: rand.String(10),
					WebSocketPath:     rand.String(10),
					ConnectionRetry:   10,
				},
				RestAPIConfig{
					OpsManagerUsername: rand.String(20),
					OpsManagerPassword: rand.String(20),
					APIPath:            rand.String(10),
				},
			},
			expectErr: true,
		},
		{
			config: &TurboCommunicationConfig{
				mediationcontainer.ServerMeta{
					TurboServer: "http:/127.0.0.1",
				},
				mediationcontainer.WebSocketConfig{
					LocalAddress: "invalid",
				},
				RestAPIConfig{
					OpsManagerUsername: rand.String(20),
					OpsManagerPassword: rand.String(20),
					APIPath:            rand.String(10),
				},
			},
			expectErr: true,
		},
		{
			config: &TurboCommunicationConfig{
				ServerMeta: mediationcontainer.ServerMeta{
					TurboServer: "http:/127.0.0.1",
				},
				RestAPIConfig: RestAPIConfig{
					OpsManagerUsername: "",
					OpsManagerPassword: rand.String(20),
					APIPath:            rand.String(10),
				},
			},
			expectErr: true,
		},
		{
			config: &TurboCommunicationConfig{
				ServerMeta: mediationcontainer.ServerMeta{
					TurboServer: "http:/127.0.0.1",
				},
				RestAPIConfig: RestAPIConfig{
					OpsManagerUsername: rand.String(20),
					OpsManagerPassword: rand.String(20),
					APIPath:            rand.String(10),
				},
			},
			expectErr: false,
		},
	}
	for _, item := range table {
		err := item.config.ValidateTurboCommunicationConfig()
		if item.expectErr {
			if err == nil {
				t.Errorf("Expects error, but go no error: %v", item.config)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		}
	}
}

func TestParseTurboCommunicationConfig(t *testing.T) {
	invalidConf := rand.String(50)
	_, err := ParseTurboCommunicationConfig(invalidConf)
	if err == nil {
		t.Error("Expects error, got no error.")
	}
}

func TestReadTurboCommunicationConfig(t *testing.T) {
	invalidConf := rand.String(50)
	_, err := readTurboCommunicationConfig(invalidConf)
	if err == nil {
		t.Error("Expects error, got no error.")
	}
}
