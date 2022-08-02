.PHONY: default server client fmt clean all release-all client-assets server-assets contributors go-bindata

BUILDTAGS=debug
default: all

fmt:
	go fmt ./...

tidy:
	go mod tidy

server: server-assets
	go build  -o bin/ -tags '$(BUILDTAGS)' cmd/ngrokd/ngrokd.go

client: client-assets
	go build -o bin/ -tags '$(BUILDTAGS)' cmd/ngrok/ngrok.go 

go-bindata:
	go install github.com/go-bindata/go-bindata/...

client-assets: go-bindata
	go-bindata -nomemcopy -pkg=assets -tags=$(BUILDTAGS) \
		-debug=$(if $(findstring debug,$(BUILDTAGS)),true,false) \
		-o=assets/client/assets/assets_$(BUILDTAGS).go -ignore=.*.go \
		assets/client/...

server-assets: go-bindata
	go-bindata -nomemcopy -pkg=assets -tags=$(BUILDTAGS) \
		-debug=$(if $(findstring debug,$(BUILDTAGS)),true,false) \
		-o=assets/server/assets/assets_$(BUILDTAGS).go -ignore=.*.go \
		assets/server/...

release-client: BUILDTAGS=release
release-client: client

release-server: BUILDTAGS=release
release-server: server

release-all: fmt release-client release-server tidy

all: fmt client server tidy

contributors:
	echo "Contributors to ngrok, both large and small:\n" > CONTRIBUTORS
	git log --raw | grep "^Author: " | sort | uniq | cut -d ' ' -f2- | sed 's/^/- /' | cut -d '<' -f1 >> CONTRIBUTORS
