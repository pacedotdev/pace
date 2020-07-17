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

release:
	cd cmd/pace && GOOS=darwin GOARCH=amd64 go build -o pace-cli-${VERSION}/macOS/pace -ldflags="-X 'main.ShortSHA=${SHORT_SHA}' -X 'main.Version=${VERSION}'"
	cd cmd/pace && GOOS=linux GOARCH=arm go build -o pace-cli-${VERSION}/linux/pace -ldflags="-X 'main.ShortSHA=${SHORT_SHA}' -X 'main.Version=${VERSION}'"
	cd cmd/pace && GOOS=windows GOARCH=amd64 go build -o pace-cli-${VERSION}/windows/pace -ldflags="-X 'main.ShortSHA=${SHORT_SHA}' -X 'main.Version=${VERSION}'"
	zip -r pace-cli-${VERSION}.zip ./cmd/pace/pace-cli-${VERSION}
	rm -rf ./cmd/pace/pace-cli-${VERSION}
	echo "Here you go... pace-cli-${VERSION}.zip"
