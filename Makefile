BIN=go-cangjie

.PHONY: dep dep-update dep-download build

all: clean test dep-download build

dep:
	go mod init
	go mod tidy

dep-update: test
	go list -m all
	go mod tidy

dep-download:
	go mod download

clean:
	go clean
	if [ -f ${BIN} ]; then rm ${BIN}; fi

test:
	go test -v ./...

build:
	go build -o ${BIN}