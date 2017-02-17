package communication

import (
	"testing"

	"github.com/turbonomic/turbo-go-sdk/util/rand"
	"reflect"
)

func TestValidateServerMeta(t *testing.T) {
	table := []struct {
		turboServer string

		expectErr bool
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
			turboServer: "invalid-url",
			expectErr:   true,
		},
	}
	for _, item := range table {
		meta := &ServerMeta{
			item.turboServer,
		}
		err := meta.validateServerMeta()
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
		err := config.validateWebSocketConfig()
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
		err := config.validRestAPIConfig()
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

func TestMediationContainerConfig_ValidateMediationContainerConfig(t *testing.T) {
	table := []struct {
		containerConfig *MediationContainerConfig
		expectErr       bool
	}{
		{
			containerConfig: &MediationContainerConfig{
				ServerMeta{
					"invalid",
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
					"http:/127.0.0.1",
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
					"http:/127.0.0.1",
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

func TestValidateTurboCommunicationConfig(t *testing.T) {
	table := []struct {
		config    *TurboCommunicationConfig
		expectErr bool
	}{
		{
			config: &TurboCommunicationConfig{
				ServerMeta{
					"invalid",
				},
				WebSocketConfig{
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
				ServerMeta{
					"http:/127.0.0.1",
				},
				WebSocketConfig{
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
				ServerMeta: ServerMeta{
					"http:/127.0.0.1",
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
				ServerMeta: ServerMeta{
					"http:/127.0.0.1",
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
		err := item.config.validateTurboCommunicationConfig()
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
	if err== nil {
		t.Error("Expects error, got no error.")
	}
}

func TestReadTurboCommunicationConfig(t *testing.T) {
	invalidConf := rand.String(50)
	_, err := readTurboCommunicationConfig(invalidConf)
	if err== nil {
		t.Error("Expects error, got no error.")
	}
}
