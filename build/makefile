.PHONY: buildandrun
BIN_FILE=cqhttp-server

build linux:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -o "${BIN_FILE}" ../cmd/main.go

build macos:
	CGO_ENABLED=0 \
	GOOS=darwin \
	GOARCH=amd64 \
	go build -o "${BIN_FILE}" ../cmd/main.go