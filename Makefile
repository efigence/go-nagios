# generate version number
version=$(shell git describe --tags --long --always|sed 's/^v//')

all:
	go build  -ldflags "-X main.version $(version)" [a-z]*go
	bash -c 'go fmt ; true'

version:
	@echo $(version)
