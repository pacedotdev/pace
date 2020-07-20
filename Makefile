githash = $(shell git rev-parse --short HEAD)
SHORT_SHA = ${githash}
VERSION = v0.1.0

.PHONY: install genoto

install:
	cd cmd/pace && go install -ldflags="-X 'main.ShortSHA=${SHORT_SHA}' -X 'main.Version=${VERSION}'"

genoto:
	oto \
		-template ./oto/client.go.plush -out ./pace.gen.go \
		-pkg pace \
		./oto
	gofmt -w ./pace.gen.go ./pace.gen.go
	oto \
		-template ./oto/cli.go.plush -out ./cmd/pace/cli.gen.go \
		-pkg pace \
		./oto
	gofmt -w ./cmd/pace/cli.gen.go ./cmd/pace/cli.gen.go

cmdbuild:
	cd cmd/pace && goreleaser --snapshot --skip-publish --rm-dist
