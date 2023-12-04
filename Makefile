BIN=congkit
DB_GENERATOR=gen-db

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

lint:
	docker run -t --rm -v ${PWD}:/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint -E revive run -v

test-concurrency:
	go test -race -run ./...

test-coverage:
	go test -cover ./...

test:
	go test -v ./...

build:
	go build -o ${DB_GENERATOR} ./cmd/db-generator
	go build -o ${BIN}

generate: build
	./gen-db