# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test
GODEP=dep

test:
	$(GOTEST) -v ./...
dep:
	$(GODEP) ensure
