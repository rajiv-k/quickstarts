BIN?=go-echo
GOOS?=linux
IMAGE_TAG?=latest

go-echo:
	CGO_ENABLED=0 GOOS=${GOOS} go build -o ${BIN} main.go

docker:
	@docker build -t go-echo:${IMAGE_TAG} -f Dockerfile.echo .

