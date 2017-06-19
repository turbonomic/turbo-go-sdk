package mediationcontainer

import (
	"reflect"
	"testing"

	"github.com/turbonomic/turbo-go-sdk/pkg/version"
	"github.com/turbonomic/turbo-go-sdk/util/rand"
)

func TestValidateServerMeta(t *testing.T) {
	table := []struct {
		turboServer string
		version     string
		expectErr   bool
	}{
		{
			turboServer: "",
			expectErr:   true,
		},
		{
			turboServer: "https://localhost:8080",
			expectErr:   false,
		},
		{
			turboServer: "https://localhost:8080",
			version:     "random",
			expectErr:   false,
		},
		{
			turboServer: "invalid-url",
			expectErr:   true,
		},
	}
	for _, item := range table {
		meta := &ServerMeta{
			TurboServer: item.turboServer,
			Version:     item.version,
		}
		err := meta.ValidateServerMeta()
		if item.expectErr {
			if err == nil {
				t.Errorf("Expected error, got no error. %s", meta)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		}
	}
}

func TestValidateServerMetaVersion(t *testing.T) {
	table := []struct {
		version string

		expectedVersion string
	}{
		{
			version:         "",
			expectedVersion: string(version.PROTOBUF_VERSION),
		},
		{
			version:         "random",
			expectedVersion: "random",
		},
	}
	for i, item := range table {
		meta := &ServerMeta{
			TurboServer: "http://localhost:8080",
			Version:     item.version,
		}
		err := meta.ValidateServerMeta()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if meta.Version != item.expectedVersion {
			t.Errorf("Test case %d failed. Expected %s, got %s", i, item.expectedVersion, meta.Version)
		}
	}
}

func TestValidateWebSocketConfig(t *testing.T) {
	table := []struct {
		localAddress    string
		wsUsername      string
		wsPassword      string
		connectionRetry int16
		wsPath          string

		expectedConfig *WebSocketConfig

		expectErr bool
	}{
		{

			expectErr: false,
		},
		{
			localAddress: "invalid_request_url",
			expectErr:    true,
		},
		{
			localAddress:    "http://1.2.3.4",
			wsUsername:      rand.String(5),
			wsPassword:      rand.String(5),
			connectionRetry: 5,
			wsPath:          rand.String(5),

			expectErr: false,
		},
	}
	for _, item := range table {
		config := &WebSocketConfig{
			LocalAddress:      item.localAddress,
			WebSocketUsername: item.wsUsername,
			WebSocketPassword: item.wsPassword,
			WebSocketPath:     item.wsPath,
			ConnectionRetry:   item.connectionRetry,
		}
		err := config.ValidateWebSocketConfig()
		if item.expectErr {
			if err == nil {
				t.Errorf("Expected error, got no error: %v", config)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
			expectedConfig := &WebSocketConfig{
				LocalAddress:      item.localAddress,
				WebSocketUsername: item.wsUsername,
				WebSocketPassword: item.wsPassword,
				WebSocketPath:     item.wsPath,
				ConnectionRetry:   item.connectionRetry,
			}
			if expectedConfig.LocalAddress == "" {
				expectedConfig.LocalAddress = defaultRemoteMediationLocalAddress
			}
			if expectedConfig.WebSocketPath == "" {
				expectedConfig.WebSocketPath = defaultRemoteMediationServer
			}
			if expectedConfig.WebSocketUsername == "" {
				expectedConfig.WebSocketUsername = defaultRemoteMediationServerUser
			}
			if expectedConfig.WebSocketPassword == "" {
				expectedConfig.WebSocketPassword = defaultRemoteMediationServerPwd
			}
			if !reflect.DeepEqual(expectedConfig, config) {
				t.Errorf("Expects %v, got %v", expectedConfig, config)
			}
		}

	}
}

func TestMediationContainerConfig_ValidateMediationContainerConfig(t *testing.T) {
	table := []struct {
		containerConfig *MediationContainerConfig
		expectErr       bool
	}{
		{
			containerConfig: &MediationContainerConfig{
				ServerMeta{
					TurboServer: "invalid",
				},
				WebSocketConfig{
					LocalAddress:      "http://127.0.0.1",
					WebSocketUsername: rand.String(10),
					WebSocketPassword: rand.String(10),
					WebSocketPath:     rand.String(10),
					ConnectionRetry:   10,
				},
			},
			expectErr: true,
		},
		{
			containerConfig: &MediationContainerConfig{
				ServerMeta{
					TurboServer: "http:/127.0.0.1",
				},
				WebSocketConfig{
					LocalAddress: "invalid",
				},
			},
			expectErr: true,
		},
		{
			containerConfig: &MediationContainerConfig{
				ServerMeta: ServerMeta{
					TurboServer: "http:/127.0.0.1",
				},
			},
			expectErr: false,
		},
	}
	for _, item := range table {
		err := item.containerConfig.ValidateMediationContainerConfig()
		if item.expectErr {
			if err == nil {
				t.Errorf("Expects error, but go no error: %v", item.containerConfig)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		}
	}
}
