package metadata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/glog"
)

type Meta struct {
	ServerAddress      string
	TargetType         string
	NameOrAddress      string
	Username           string
	TargetIdentifier   string
	Password           string
	LocalAddress       string
	WebSocketUsername  string
	WebSocketPassword  string
	OpsManagerUsername string
	OpsManagerPassword string
}

// Create a new Meta from file.
func NewMetaFromFile(metaConfigFilePath string) (*Meta, error) {
	glog.V(4).Infof("Now read configration from %s", metaConfigFilePath)
	metaConfig := readConfig(metaConfigFilePath)
	return NewMeta(metaConfig.ServerAddress, metaConfig.TargetType, metaConfig.NameOrAddress, metaConfig.TargetIdentifier, metaConfig.Username, metaConfig.Password, metaConfig.LocalAddress, metaConfig.WebSocketUsername, metaConfig.Password, metaConfig.OpsManagerUsername, metaConfig.OpsManagerPassword)

}

func NewMeta(serverAddr, targetType, nameOrAddress, targetID, usrn, passd, localAddr, wsUsrn, wsPassd, opsManUsrn, opsManPassd string) (*Meta, error) {
	if len(serverAddr) == 0 {
		return nil, fmt.Errorf("Error getting server address.")
	}
	glog.V(2).Infof("VMTurbo Server Address is %s", serverAddr)

	if len(targetID) == 0 {
		return nil, fmt.Errorf("Error gettring target identifier.")
	}
	glog.V(3).Infof("TargetIdentifier is %s", targetID)

	if len(nameOrAddress) == 0 {
		return nil, fmt.Errorf("Error getting NameOrAddress for exmaple probe.")
	}

	if len(usrn) == 0 {
		return nil, fmt.Errorf("Erorr getting username for target configuration.")
	}

	if len(targetType) == 0 {
		return nil, fmt.Errorf("Erorr getting target type for target configuration.")
	}

	if len(passd) == 0 {
		return nil, fmt.Errorf("Erorr getting password for target configuration.")
	}

	if len(localAddr) == 0 {
		return nil, fmt.Errorf("Erorr getting local address from meata data.")
	}

	if len(opsManUsrn) == 0 {
		return nil, fmt.Errorf("Erorr getting OpsManagerUsername from meta data.")
	}

	if len(opsManPassd) == 0 {
		return nil, fmt.Errorf("Erorr getting OpsManagerPassword from meta data.")
	}
	meta := &Meta{
		ServerAddress:      serverAddr,
		TargetType:         targetType,
		NameOrAddress:      nameOrAddress,
		Username:           usrn,
		TargetIdentifier:   targetID,
		Password:           passd,
		LocalAddress:       localAddr,
		WebSocketUsername:  wsUsrn,
		WebSocketPassword:  wsPassd,
		OpsManagerUsername: opsManUsrn,
		OpsManagerPassword: opsManPassd,
	}
	return meta, nil
}

// Get the config from file.
func readConfig(path string) Meta {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		glog.Errorf("File error: %v\n", e)
		os.Exit(1)
	}
	var metaData Meta
	json.Unmarshal(file, &metaData)
	glog.V(4).Infof("Results: %v\n", metaData)
	return metaData
}
