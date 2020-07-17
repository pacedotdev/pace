githash = $(shell git rev-parse --short HEAD)
SHORT_SHA = ${githash}
VERSION = v0.1.0

.PHONY: install genoto

install:
	cd cmd/pace && go install

genoto:
	oto \
		-template ./oto/templates/client.go.plush -out ./pace.gen.go \
		-pkg pace \
		./oto/definition
	gofmt -w ./pace.gen.go ./pace.gen.go
	oto \
		-template ./oto/templates/cli.go.plush -out ./cmd/pace/cli.gen.go \
		-pkg pace \
		./oto/definition
	gofmt -w ./cmd/pace/cli.gen.go ./cmd/pace/cli.gen.go

release:
	cd cmd/pace && go build -o pace -ldflags="-X 'main.ShortSHA=$SHORT_SHA'"
