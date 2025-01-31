EXECUTABLE=smtool
VERSION=$(shell git describe --tags --always --long --dirty)
WINDOWS=$(EXECUTABLE)_windows_amd64_$(VERSION)
LINUX=$(EXECUTABLE)_linux_amd64_$(VERSION)
DARWIN=$(EXECUTABLE)_darwin_amd64_$(VERSION)

.PHONY: all test clean

all: test build

test:
	mkdir coverage -p && go test -v -coverprofile coverage/cover.out && go tool cover -html coverage/cover.out -o coverage/cover.html

build: windows linux darwin
	@echo version: $(VERSION)

windows: $(WINDOWS)

linux: $(LINUX)

darwin: $(DARWIN)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o bin/win/$(EXECUTABLE).exe -ldflags="-s -w -X main.version=$(VERSION)" && zip --junk-paths bin/$(WINDOWS).zip bin/win/$(EXECUTABLE).exe

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o bin/linux/$(EXECUTABLE) -ldflags="-s -w -X main.version=$(VERSION)" && zip --junk-paths bin/$(LINUX).zip bin/linux/$(EXECUTABLE)

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -v -o bin/darwin/$(EXECUTABLE) -ldflags="-s -w -X main.version=$(VERSION)" && zip --junk-paths bin/$(DARWIN).zip bin/darwin/$(EXECUTABLE)

clean:
	rm -rf bin