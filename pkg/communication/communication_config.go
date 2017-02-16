package communication

import (
	"errors"
	"fmt"

	"encoding/json"
	"github.com/golang/glog"
	"io/ioutil"
	"net/url"
)

const (
	defaultRemoteMediationServer       string = "/vmturbo/remoteMediation"
	defaultRemoteMediationServerUser   string = "vmtRemoteMediation"
	defaultRemoteMediationServerPwd    string = "vmtRemoteMediation"
	defaultRemoteMediationLocalAddress string = "http://127.0.0.1"
)

type ServerMeta struct {
	TurboServer string `json:"turboServer,omitempty"`
}

func (meta *ServerMeta) validateServerMeta() error {
	if meta.TurboServer == "" {
		return errors.New("Turbo Server URL is missing")
	}
	if _, err := url.Parse(meta.TurboServer); err != nil {
		return fmt.Errorf("Invalid turbo address url: %v", meta)
	}
	return nil
}

type WebSocketConfig struct {
	LocalAddress      string `json:"localAddress,omitempty"`
	WebSocketUsername string `json:"websocketUsername,omitempty"`
	WebSocketPassword string `json:"websocketPassword,omitempty"`
	ConnectionRetry   int16  `json:"connectionRetry,omitempty"`
	WebSocketPath     string `json:"websocketPath,omitempty"`
}

func (wsc *WebSocketConfig) validateWebSocketConfig() error {
	if wsc.LocalAddress == "" {
		wsc.LocalAddress = defaultRemoteMediationLocalAddress
	}
	// Make sure the local address string provided is a valid URL
	if _, err := url.Parse(wsc.LocalAddress); err != nil {
		return fmt.Errorf("Invalid local address url found in WebSocket config: %v", wsc)
	}

	if wsc.WebSocketPath == "" {
		wsc.WebSocketPath = defaultRemoteMediationServer
	}
	if wsc.WebSocketUsername == "" {
		wsc.WebSocketUsername = defaultRemoteMediationServerUser
	}
	if wsc.WebSocketPassword == "" {
		wsc.WebSocketPassword = defaultRemoteMediationServerPwd
	}
	return nil
}

type RestAPIConfig struct {
	OpsManagerUserName string `json:"opsManagerUsername,omitempty"`
	OpsManagerPassword string `json:"opsManagerPassword,omitempty"`
	APIPath            string `json:"apiPath,omitempty"`
}

func (rc *RestAPIConfig) validRestAPIConfig() error {
	if rc.OpsManagerUserName == "" || rc.OpsManagerPassword == "" {
		return errors.New("Either username or password for API is not provided.")
	}
	return nil
}

type MediationContainerConfig struct {
	ServerMeta
	WebSocketConfig
}

// Validate the mediation container config and set default value if necessary.
func (containerConfig *MediationContainerConfig) ValidateMediationContainerConfig() error {
	if err := containerConfig.validateServerMeta(); err != nil {
		return err
	}
	if err := containerConfig.validateWebSocketConfig(); err != nil {
		return err
	}
	glog.V(4).Infof("The mediation container config is %v", containerConfig)
	return nil
}

// Configuration parameters for communicating with the Turbo server
type TurboCommunicationConfig struct {
	ServerMeta      `json:"serverMeta,omitempty"`
	WebSocketConfig `json:"websocketConfig,omitempty"`
	RestAPIConfig   `json:"restAPIConfig,omitempty"`
}

func ParseTurboCommunicationConfig(configFile string) (*TurboCommunicationConfig, error) {
	// load the config
	turboCommConfig, err := readTurboCommunicationConfig(configFile)
	if turboCommConfig == nil {
		return nil, err
	}
	glog.V(3).Infof("TurboCommunicationConfig Config: %v", turboCommConfig)

	// validate the config
	if err := turboCommConfig.validateServerMeta(); err != nil {
		return nil, err
	}
	if err := turboCommConfig.validateWebSocketConfig(); err != nil {
		return nil, err
	}
	if err := turboCommConfig.validRestAPIConfig(); err != nil {
		return nil, err
	}
	glog.V(3).Infof("Loaded Turbo Communication Config: %v", turboCommConfig)
	return turboCommConfig, nil
}

func readTurboCommunicationConfig(path string) (*TurboCommunicationConfig, error) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, fmt.Errorf("File error: %v\n" + e.Error())
	}
	var config TurboCommunicationConfig
	err := json.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("Unmarshall error :%v", err.Error())
	}
	return &config, nil
}
