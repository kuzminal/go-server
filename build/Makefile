build-docker-image:
	docker build -t http-server ../

all: install generate run

install:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0

generate:
	oapi-codegen -config ../configs/server.cfg.yaml ../api/hello.yaml

run:
	go run server.go
