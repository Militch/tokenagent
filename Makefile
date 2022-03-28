PROJECTNAME := tokenagent
APPNAME := tokenagent

ifeq ($(OS),Windows_NT)
	BINARY_EXT := .exe
else
	BINARY_EXT :=
endif

BINARY := $(APPNAME)$(BINARY_EXT)

.PHONY: all
all: build

.PHONY: build
build: 
	go build -o $(BINARY) main.go

devtools:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0 1>/dev/null

precheck: devtools
	@${GOPATH}/bin/golangci-lint run 1>/dev/null

test: precheck
	@go test ./...
	@echo "Test successful"

clean:
	$(if $(filter $(OS),Windows_NT), del ${BINARY}, rm -f ${BINARY})