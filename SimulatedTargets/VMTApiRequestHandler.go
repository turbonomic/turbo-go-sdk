package main

import (
	"github.com/golang/glog"
	"net/http"
)

// This struct holds the authorization information and address for connecting to VMTurbo API
type VMTApiRequestHandler struct {
	vmtServerAddr      string
	opsManagerUsername string
	opsManagerPassword string
}

// this helper function servers to send REST api calls to the VMTServer using opsmanager authentication
func (vmtapi *VMTApiRequestHandler) vmtApiPost(postPath, requestStr string) (*http.Response, error) {
	fullUrl := "http://" + vmtapi.vmtServerAddr + "/vmturbo/api" + postPath + requestStr
	req, err := http.NewRequest("POST", fullUrl, nil)
	req.SetBasicAuth(vmtapi.opsManagerUsername, vmtapi.opsManagerPassword)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		glog.Infof("Log: error getting response")
		return nil, err
	}
	defer response.Body.Close()
	return response, nil
}
