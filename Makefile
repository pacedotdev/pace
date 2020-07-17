
.PHONY: install genoto

install:
	cd cmd/pace && go install

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
