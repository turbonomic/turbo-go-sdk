package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/turbonomic/turbo-go-sdk/example/sdk-example-probe/pkg/turboapi/api"

	"github.com/golang/glog"
)

type Request struct {
	verb    string
	baseURL *url.URL

	pathPrefix string
	params     url.Values

	basicAuth *BasicAuthentication

	resource     api.ResourceType
	resourceName string

	err error
}

type Result struct {
	body []byte
	err  error
}

func NewRequest(verb string, baseURL *url.URL, apiPath string) *Request {
	if len(apiPath) != 0 && !strings.HasPrefix(apiPath, "/") {
		apiPath = path.Join("/", apiPath)
	}
	return &Request{
		verb:       verb,
		baseURL:    baseURL,
		pathPrefix: apiPath,
	}
}

func (this *Request) BasicAuthentication(basicAuth *BasicAuthentication) *Request {
	this.basicAuth = basicAuth
	return this
}

// Set the kind of the api resource that the request is made to.
func (this *Request) Resource(resource api.ResourceType) *Request {
	if this.err != nil {
		return this
	}

	if this.resource != "" {
		this.err = fmt.Errorf("Resource has already been set. Cannot be changed!")
		return this
	}
	this.resource = resource
	return this
}

func (this *Request) Name(resourceName string) *Request {
	if this.err != nil {
		return this
	}

	if this.resourceName != "" {
		this.err = fmt.Errorf("Resource name has already been set to %s. Cannot be changed!", this.resourceName)
		return this
	}
	if len(resourceName) == 0 {
		this.err = fmt.Errorf("Resource name cannot be empty.")
		return this
	}
	this.resourceName = resourceName
	return this
}

// Set parameters for the request.
func (this *Request) Param(paramName, value string) *Request {
	if this.params == nil {
		this.params = make(url.Values)
	}
	this.params[paramName] = append(this.params[paramName], value)

	return this
}

// URL returns the current working URL.
func (this *Request) URL() *url.URL {
	p := this.pathPrefix

	if len(this.resource) != 0 {
		p = path.Join(p, strings.ToLower(string(this.resource)))
	}

	if len(this.resourceName) != 0 {
		p = path.Join(p, this.resourceName)
	}

	finalURL := &url.URL{}
	if this.baseURL != nil {
		*finalURL = *this.baseURL
	}
	finalURL.Path = p

	query := url.Values{}
	for key, values := range this.params {
		for _, value := range values {
			query.Add(key, value)
		}
	}

	finalURL.RawQuery = query.Encode()

	return finalURL
}

func (this *Request) Do() (string, error) {
	var result Result
	err := this.request(func(resp *http.Response) {
		result = parseAPICallResponse(resp)
	})
	if err != nil {
		return "", err
	}
	if result.err != nil {
		return "", err
	}
	return string(result.body), nil
}

func (this *Request) request(fn func(*http.Response)) error {
	if this.err != nil {
		return this.err
	}

	client := http.DefaultClient

	url := this.URL().String()
	glog.V(4).Infof("The request url is %s", url)
	req, err := http.NewRequest(this.verb, url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(this.basicAuth.username, this.basicAuth.password)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	fn(resp)
	return nil
}

func parseAPICallResponse(resp *http.Response) Result {
	if resp == nil {
		return Result{
			body: nil,
			err:  fmt.Errorf("response sent in is nil"),
		}
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Result{
			body: nil,
			err:  fmt.Errorf("Error reading response body:%++v", err),
		}
	}

	return Result{
		body: content,
		err:  nil,
	}
}
