.PHONY:  testall test testl testv coverage threshold lint run depgraph install_admin server install_air install_mockery install_overcover install_libs
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run .
GOCOV=$(GOCMD) tool cover -html=coverage.filtered.out
GOTEST=$(GOCMD) test -tags test -short
GOGET=$(GOCMD) get
GODEP=godepgraph -s -o  github.com/webmalc/vishleva-backend github.com/webmalc/vishleva-backend | dot -Tpng -o godepgraph.png
BINARY_NAME=vishleva_backend.app

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	GOENV=test $(GOTEST) ./... -coverprofile=coverage.out
	cat coverage.out | grep -v "mock" | grep -v "bindatafs" > coverage.filtered.out

testv:
	GOENV=test $(GOTEST) -v ./... -coverprofile=coverage.out
	cat coverage.out | grep -v "mock" | grep -v "bindatafs" > coverage.filtered.out

depgraph:
	$(GODEP)

coverage:
	$(GOCOV)

threshold:
	overcover --coverprofile coverage.filtered.out --threshold 80 --summary
testl: testv lint

testall: test lint threshold

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

install_admin:
	$(GOCMD) get -u github.com/qor/bindatafs/...
	$(GOCMD) get -u github.com/qor/admin

lint:
	golangci-lint run ./...

run:
	$(GORUN) $(c)

install_air:
	go install github.com/air-verse/air@latest
	air init

# mockery --dir=common/messenger --name=VkAPISender --filename=interfaces.go --output=common/messenger/mocks --outpkg=mocks --filename=vk_api_sender.go
install_mockery:
	go install github.com/vektra/mockery/v2@v2.50.0

install_overcover:
	go install github.com/klmitch/overcover@latest

install_libs: install_mockery install_overcover install_air


server:
	air server
