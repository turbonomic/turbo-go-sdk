package mediationcontainer

import (
	"github.com/stretchr/testify/assert"
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
		expectedConfig  *WebSocketConfig
		expectErr       bool
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
			expectErr:       false,
		},
	}
	for _, item := range table {
		config := &WebSocketConfig{
			LocalAddress:      item.localAddress,
			WebSocketUsername: item.wsUsername,
			WebSocketPassword: item.wsPassword,
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
				ConnectionRetry:   item.connectionRetry,
			}
			if expectedConfig.LocalAddress == "" {
				expectedConfig.LocalAddress = defaultRemoteMediationLocalAddress
			}
			expectedConfig.WebSocketEndpoints = defaultRemoteMediationServerEndpoints
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
					ConnectionRetry:   10,
				},
				"",
				SdkProtocolConfig{
					RegistrationTimeoutSec:       60,
					RestartOnRegistrationTimeout: false,
				},
			},
			expectErr: true,
		},
		{
			containerConfig: &MediationContainerConfig{
				ServerMeta{
					TurboServer: "invalid",
				},
				WebSocketConfig{
					LocalAddress:      "http://127.0.0.1",
					WebSocketUsername: rand.String(10),
					WebSocketPassword: rand.String(10),
					ConnectionRetry:   10,
				},
				"",
				SdkProtocolConfig{
					RegistrationTimeoutSec:       60,
					RestartOnRegistrationTimeout: true,
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
				"foo",
				SdkProtocolConfig{
					RegistrationTimeoutSec:       60,
					RestartOnRegistrationTimeout: false,
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

func TestValidateSdkProtocolConfig(t *testing.T) {
	table := []struct {
		config   *SdkProtocolConfig
		expected *SdkProtocolConfig
	}{
		{
			config: &SdkProtocolConfig{},
			expected: &SdkProtocolConfig{
				RegistrationTimeoutSec:       DefaultRegistrationTimeOut,
				RestartOnRegistrationTimeout: false,
			},
		},
		{
			config: &SdkProtocolConfig{
				RegistrationTimeoutSec:       30,
				RestartOnRegistrationTimeout: true,
			},
			expected: &SdkProtocolConfig{
				RegistrationTimeoutSec:       DefaultRegistrationTimeOut,
				RestartOnRegistrationTimeout: true,
			},
		},
		{
			config: &SdkProtocolConfig{
				RegistrationTimeoutSec: 600,
			},
			expected: &SdkProtocolConfig{
				RegistrationTimeoutSec:       600,
				RestartOnRegistrationTimeout: false,
			},
		},
	}
	for _, item := range table {
		item.config.ValidateSdkProtocolConfig()

		if item.expected != nil {
			assert.Equal(t, item.expected.RegistrationTimeoutSec, item.config.RegistrationTimeoutSec)
			assert.Equal(t, item.expected.RestartOnRegistrationTimeout, item.config.RestartOnRegistrationTimeout)
		}
	}
}
