package vmtapi

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

// api information for requests to VMT server
type VmtApi struct {
	VmtUrl    string
	Username  string
	Password  string
	extConfig map[string]string
}

func NewVmtApi(vmtUrl string, username string, pwd string) *VmtApi {
	extCongfix := make(map[string]string)
	extCongfix["Username"] = username
	extCongfix["Password"] = pwd
	return &VmtApi{
		VmtUrl:    vmtUrl,
		extConfig: extCongfix,
	}
}

func (vmtApi *VmtApi) Post(postUrl, requestDataString string) (string, error) {
	return vmtApi.apiPost(postUrl, requestDataString)
}

func (vmtApi *VmtApi) Get(getUrl string) (string, error) {
	return vmtApi.apiGet(getUrl)
}

func (vmtApi *VmtApi) Delete(getUrl string) (string, error) {
	return vmtApi.apiDelete(getUrl)
}

// Call vmturbo api. return response
func (vmtApi *VmtApi) apiPost(postUrl, requestDataString string) (string, error) {
	fullUrl := "http://" + vmtApi.VmtUrl + "/vmturbo/api" + postUrl + requestDataString
	fmt.Println("The full Url is ", fullUrl)
	req, err := http.NewRequest("POST", fullUrl, nil)

	req.SetBasicAuth(vmtApi.extConfig["Username"], vmtApi.extConfig["Password"])
	fmt.Println(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting response: %s", err)
		return "", err
	}

	respContent, err := parseAPICallResponse(resp)
	if err != nil {
		fmt.Println("Error getting response: %s", err)
		return "", err
	}
	fmt.Println("Post Succeed: %s", string(respContent))

	defer resp.Body.Close()
	return respContent, nil
}

// Call vmturbo api. return response
func (vmtApi *VmtApi) apiGet(getUrl string) (string, error) {
	fullUrl := "http://" + vmtApi.VmtUrl + "/vmturbo/api" + getUrl
	glog.V(4).Info("The full Url is ", fullUrl)
	req, err := http.NewRequest("GET", fullUrl, nil)

	req.SetBasicAuth(vmtApi.extConfig["Username"], vmtApi.extConfig["Password"])
	glog.V(4).Info(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Errorf("Error getting response: %s", err)
		return "", err
	}
	respContent, err := parseAPICallResponse(resp)
	if err != nil {
		glog.Errorf("Error getting response: %s", err)
		return "", err
	}
	glog.V(4).Infof("Get Succeed: %s", string(respContent))
	defer resp.Body.Close()
	return respContent, nil
}

// Delete API call
func (vmtApi *VmtApi) apiDelete(getUrl string) (string, error) {
	fullUrl := "http://" + vmtApi.VmtUrl + "/vmturbo/api" + getUrl
	glog.V(4).Info("The full Url is ", fullUrl)
	req, err := http.NewRequest("DELETE", fullUrl, nil)

	req.SetBasicAuth(vmtApi.extConfig["Username"], vmtApi.extConfig["Password"])
	glog.V(4).Info(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Errorf("Error getting response: %s", err)
		return "", err
	}
	respContent, err := parseAPICallResponse(resp)
	if err != nil {
		glog.Errorf("Error getting response: %s", err)
		return "", err
	}
	glog.V(4).Infof("DELETE call Succeed: %s", string(respContent))
	defer resp.Body.Close()
	return respContent, nil
}

// this method takes in a reservation response and should return the reservation uuid, if there is any
func parseAPICallResponse(resp *http.Response) (string, error) {
	if resp == nil {
		glog.V(4).Infof("response from VMTServer for Reservation UUID is nil")
		return "", errors.New("response sent in is nil")
	}
	glog.V(4).Infof("response body is %s", resp.Body)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorf("Error after ioutil.ReadAll: %s", err)
		return "", err
	}
	glog.V(4).Infof("response content is %s", string(content))

	return string(content), nil
}
