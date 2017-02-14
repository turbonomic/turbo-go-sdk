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

func (cb *ConfigBuilder) APIPath(apiPath string) *ConfigBuilder {
	cb.apiPath = apiPath
	return cb
}

func (cb *ConfigBuilder) BasicAuthentication(usrn, passd string) *ConfigBuilder {
	cb.basicAuth = &BasicAuthentication{
		username: usrn,
		password: passd,
	}
	return cb
}

func (cb *ConfigBuilder) Create() *Config {
	return &Config{
		ServerAddress: cb.serverAddress,
		// If API path not specified, use the default API path.
		APIPath: func() string {
			if cb.apiPath == "" {
				return defaultAPIPath
			}
			return cb.apiPath
		}(),
		BasicAuth: cb.basicAuth,
	}
}
