module github.com/turbonomic/turbo-go-sdk

go 1.13

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.1
	github.com/gorilla/websocket v1.2.0
	github.com/stretchr/testify v1.3.0
	github.com/turbonomic/turbo-api v0.0.0-20180816193551-ed948ba97e70
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a
)

replace github.com/turbonomic/turbo-api => ../turbo-api
