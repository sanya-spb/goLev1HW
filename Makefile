PROJECT?=github.com/sanya-spb/goLev1HW

GOOS?=linux
GOARCH?=amd64

RELEASE := $(shell git tag -l | tail -1 | grep -E "v.+"|| echo devel)
COMMIT := git-$(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
COPYRIGHT := "sanya-spb"

build:
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 go build \
		-ldflags "-s -w -X ${PROJECT}/utils/version.version=${RELEASE} \
		-X ${PROJECT}/utils/version.commit=${COMMIT} \
		-X ${PROJECT}/utils/version.buildTime=${BUILD_TIME} \
		-X ${PROJECT}/utils/version.copyright=${COPYRIGHT}" \
		-o app_main 