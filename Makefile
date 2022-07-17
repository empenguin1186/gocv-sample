.ONESHELL:
.PHONY: docker codegen

GOCV_VERSION?="v0.31.0"
OPENCV_VERSION?=4.6.0
GOVERSION?=1.16.2

docker:
	docker build --build-arg OPENCV_VERSION=$(OPENCV_VERSION) --build-arg GOVERSION=$(GOVERSION) .

codegen:
	docker run --rm -v "$(CURDIR):/app" openapitools/openapi-generator-cli generate \
	-i "/app/spec/openapi.yaml" \
	-g go-server \
	-o /app/code_generated