SOURCE_DIRS = pkg
PACKAGES := go list ./... | grep -v /vendor | grep -v /out

build:
	go build ./...

test:
	@go test -v -race ./pkg/...

.PHONY: fmtcheck
fmtcheck:
	@gofmt -l $(SOURCE_DIRS) | grep ".*\.go" | grep -v "pkg/version/version.pb.go"; if [ "$$?" = "0" ]; then exit 1; fi

.PHONY: vet
vet:
	@go vet $(shell $(PACKAGES))
