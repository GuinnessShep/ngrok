.PHONY: default server client fmt clean all release-all client-assets server-assets contributors go-bindata

BUILDTAGS=debug
GO111MODULE=on
default: all

fmt:
	go fmt ./...

tidy:
	go mod tidy

go-bindata: export GOOS=
go-bindata: export GOARCH=
go-bindata:
	go get github.com/go-bindata/go-bindata
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

%-server: export BUILDTAGS=$*
%-server: server-assets
	go build -o bin/$(if $(GOOS),$(GOOS)_${GOARCH},"")/ \
		-tags '$(BUILDTAGS)' cmd/ngrokd/ngrokd.go

%-client: export BUILDTAGS=$*
%-client: client-assets
	go build -o bin/$(if $(GOOS),$(GOOS)_${GOARCH},"")/ \
		-tags '$(BUILDTAGS)' cmd/ngrok/ngrok.go

all: fmt client server tidy

contributors:
	echo "Contributors to ngrok, both large and small:\n" > CONTRIBUTORS
	git log --raw | grep "^Author: " | sort | uniq | cut -d ' ' -f2- | sed 's/^/- /' | cut -d '<' -f1 >> CONTRIBUTORS
