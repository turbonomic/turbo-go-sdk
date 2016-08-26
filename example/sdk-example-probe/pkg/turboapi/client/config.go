package client

import (
	"net/url"
)

const (
	defaultAPIPath = "vmturbo/api/"
)

type Config struct {
	// ServerAddress is  a string, which must be a host:port pair.
	ServerAddress *url.URL
	// APIPath is a sub-path that points to an API root.
	APIPath string

	// For Basic authentication
	BasicAuth *BasicAuthentication
}

type ConfigBuilder struct {
	serverAddress *url.URL
	apiPath       string
	basicAuth     *BasicAuthentication
}

func NewConfigBuilder(serverAddress *url.URL) *ConfigBuilder {
	return &ConfigBuilder{
		serverAddress: serverAddress,
	}
}

func (this *ConfigBuilder) APIPath(apiPath string) *ConfigBuilder {
	this.apiPath = apiPath
	return this
}

func (this *ConfigBuilder) BasicAuthentication(usrn, passd string) *ConfigBuilder {
	this.basicAuth = &BasicAuthentication{
		username: usrn,
		password: passd,
	}
	return this
}

func (this *ConfigBuilder) Create() *Config {
	return &Config{
		ServerAddress: this.serverAddress,
		// If API path not specified, use the default API path.
		APIPath: func() string {
			if this.apiPath == "" {
				return defaultAPIPath
			}
			return this.apiPath
		}(),
		BasicAuth: this.basicAuth,
	}
}
