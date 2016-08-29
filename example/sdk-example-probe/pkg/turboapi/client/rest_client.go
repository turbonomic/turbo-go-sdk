package client

import (
	"net/url"
	"strings"
)

type RESTClient struct {
	baseURL *url.URL
	apiPath string

	basicAuth *BasicAuthentication
}

func NewRESTClient(baseURL *url.URL, apiPath string, ba *BasicAuthentication) *RESTClient {
	base := *baseURL
	if !strings.HasSuffix(base.Path, "/") {
		//		base.Path += "/"
	}

	return &RESTClient{
		baseURL: &base,
		apiPath: apiPath,

		basicAuth: ba,
	}
}

// NOTE: Currently use basic authentication.
func (this *RESTClient) Verb(verb string) *Request {
	request := NewRequest(verb, this.baseURL, this.apiPath).BasicAuthentication(this.basicAuth)
	return request
}

func (this *RESTClient) Get() *Request {
	return this.Verb("GET")
}

func (this *RESTClient) Post() *Request {
	return this.Verb("POST")
}

func (this *RESTClient) Delete() *Request {
	return this.Verb("DELETE")
}
