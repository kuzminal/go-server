LOCAL_BIN:=$(CURDIR)/../bin

build-docker-image:
	docker build -t http-server ../

all: install generate run

install:
	GOBIN=$(LOCAL_BIN) go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0

generate:
	../bin/oapi-codegen -config ../configs/server.cfg.yaml ../api/hello.yaml

run:
	go run ../cmd/server -conf ../configs/local-config.yaml
