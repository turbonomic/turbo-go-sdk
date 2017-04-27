package mediationcontainer

import (
	"testing"
	"reflect"
)

func TestCreateClientWebSocketTransport(t *testing.T) {
	table := []struct {
		localAddress            string
		turboServer             string
		expectsWebSocketAddress string

		expectsErr bool
	}{
		{
			turboServer:             "https://127.0.0.1",
			localAddress:            "http://127.0.0.1",
			expectsWebSocketAddress: "wss://127.0.0.1",

			expectsErr: false,
		},
		{
			turboServer:             "http://127.0.0.1",
			localAddress:            "http://127.0.0.1",
			expectsWebSocketAddress: "ws://127.0.0.1",

			expectsErr: false,
		},
		{
			turboServer:             "http://127.0.0.1",
			localAddress:            "invalid",
			expectsWebSocketAddress: "ws://127.0.0.1",

			expectsErr: true,
		},
		{
			turboServer: "invalid",

			expectsErr: true,
		},
	}

	for _, item := range table {
		containerConfig := &MediationContainerConfig{
			ServerMeta{
				TurboServer: item.turboServer,
			},
			WebSocketConfig{
				LocalAddress: item.localAddress,
			},
		}
		wsConfig, err := CreateWebSocketConnectionConfig(containerConfig)
		if item.expectsErr {
			if err == nil {
				t.Error("Expects error, got no error.")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			expectedWebSocketConfig := &WebSocketConnectionConfig{
				ServerMeta{
					TurboServer: item.expectsWebSocketAddress,
				},
				WebSocketConfig{
					LocalAddress: item.localAddress,
				},
			}
			if !reflect.DeepEqual(expectedWebSocketConfig, wsConfig) {
				t.Errorf("\nExpect %v,\n got   %v", expectedWebSocketConfig, wsConfig)
			}
		}
	}
}

// Test the status changes after function call.
func TestCloseAndResetWebSocket(t *testing.T) {
	table := []struct {
		currStatus TransportStatus
	}{
		{
			currStatus:Closed,
		},
		{
			currStatus:Ready,
		},
	}

	for _, item:= range table {
		// A fake WebSocket transport only for testing status change.
		fakeWebSocketTransport := &ClientWebSocketTransport{}
		fakeWebSocketTransport.status = item.currStatus
		fakeWebSocketTransport.closeAndResetWebSocket()
		if fakeWebSocketTransport.status != Closed {
			t.Errorf("Expected status is %s, got %s", Closed, fakeWebSocketTransport.status)
		}
	}
}