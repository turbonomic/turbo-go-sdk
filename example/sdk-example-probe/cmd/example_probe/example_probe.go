package main

import (
	"flag"
	"fmt"
	"net/url"
	"time"

	"github.com/turbonomic/turbo-go-sdk/example/sdk-example-probe/pkg/communication"
	"github.com/turbonomic/turbo-go-sdk/example/sdk-example-probe/pkg/metadata"
	"github.com/turbonomic/turbo-go-sdk/example/sdk-example-probe/pkg/probe"
	"github.com/turbonomic/turbo-go-sdk/example/sdk-example-probe/pkg/turboapi/client"

	"github.com/golang/glog"
)

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	flag.Parse()

	serverAddress := "192.168.1.11:8080"
	targetType := "ExampleProbe"
	targetIdentifier := "example-probe"
	nameOrAddress := "example"
	username := "administrator"
	password := "fake_passd"
	localAddress := "http://127.0.0.1"
	opsManUsrn := "administrator"
	opsManPassd := "a"

	glog.V(3).Infof("Building metadata")

	meta, err := metadata.NewMeta(
		serverAddress, targetType, nameOrAddress, targetIdentifier, username, password, localAddress, "", "", opsManUsrn, opsManPassd)
	if err != nil {
		glog.Fatal("Cannot create meta data from the given info.")
	}

	topologyAccessor, err := probe.NewTopologyGenerator(2, 3)
	if err != nil {
		glog.Fatal("Error getting topology accessor: %v", err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)
	go func() {
		for {
			select {
			case <-stopCh:
				return
			default:
			}

			topologyAccessor.UpdateResource()

			t := time.NewTimer(time.Minute * 1)
			select {
			case <-stopCh:
				return
			case <-t.C:
			}
		}

	}()

	serverMsgHandler, err := createHandler(meta, topologyAccessor)
	if err != nil {
		glog.Fatal("Cannot create server message handler from the given info.")
	}

	communicator := communication.NewCommunicator(meta, serverMsgHandler)
	communicator.Start()
	err = communicator.RegisterExampleProbe(targetType)
	if err != nil {
		glog.Fatal("Register Example Probe failed: %v", err)
	}
}

func createHandler(meta *metadata.Meta, topoAccessor *probe.TopologyGenerator) (*communication.ExampleServerMessageHandler, error) {
	turboAPIClient, err := createTurboAPIClientFromMeta(meta)

	if err != nil {
		return nil, err
	}

	handler := communication.NewExampleServerMessageHandler(turboAPIClient, meta, topoAccessor)
	return handler, nil
}

func createTurboAPIClientFromMeta(meta *metadata.Meta) (*client.Client, error) {
	serverAddress, err := url.Parse("http://" + meta.ServerAddress)
	if err != nil {
		return nil, fmt.Errorf("Invalid server address: %v", err)
	}
	apiClientConfig := client.NewConfigBuilder(serverAddress).BasicAuthentication(meta.OpsManagerUsername, meta.OpsManagerPassword).Create()
	return client.NewAPIClient(apiClientConfig), nil
}
