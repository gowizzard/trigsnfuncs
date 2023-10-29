# Here you can reformat, check or build the binary
BINARY_NAME=tfusion
APP_PATH=cmd/${BINARY_NAME}/main.go

fmt:
	@go fmt ./...

vet:
	@go vet ./...

tidy:
	@go mod tidy

download:
	@go mod download

test:
	@go test -v ./... -bench=.

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

doc:
	@godoc -play=true -goroot=/usr/local/go -http=:6060

run:
	@go run ${APP_PATH}

build:
	go mod download
	go build ${LDFLAGS} -o ${BINARY_NAME} ${APP_PATH}

apply:
	@terraform -chdir=infrastructure apply -auto-approve