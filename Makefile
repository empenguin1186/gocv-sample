.ONESHELL:
.PHONY: docker codegen zip-package

GOCV_VERSION?="v0.31.0"
OPENCV_VERSION?=4.6.0
GOVERSION?=1.16.2

docker:
	docker build -t gocv-sample --build-arg OPENCV_VERSION=$(OPENCV_VERSION) --build-arg GOVERSION=$(GOVERSION) .

codegen:
	docker run --rm -v "$(CURDIR):/app" openapitools/openapi-generator-cli generate \
	-i "/app/docs/openapi.yaml" \
	-g go-server \
	-o /app/code_generated

zip-package:
	GOOS=linux CGO_ENABLED=0 go build cmd/main.go
	chmod 755 main
	zip function.zip main