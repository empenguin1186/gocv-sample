FROM gocv/opencv:4.6.0 AS build-env

ENV GOPATH /go

COPY . /go/src/gocv-sample/

WORKDIR /go/src/gocv-sample
RUN go build -tags example -o /build/server ./cmd/

WORKDIR /go/src/
RUN go install github.com/go-delve/delve/cmd/dlv@latest

FROM gocv/opencv:4.6.0

WORKDIR /
COPY --from=build-env /build/server /
COPY --from=build-env /go/src/gocv-sample/data/haarcascade_frontalface_default.xml /data/
COPY --from=build-env /go/bin/dlv /

EXPOSE 8080 40000

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/server"]