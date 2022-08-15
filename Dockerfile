FROM gocv/opencv:4.6.0

ENV GOPATH /go

COPY . /go/src/gocv-sample/

WORKDIR /go/src/gocv-sample
RUN go build -tags example -o /build/main ./cmd/server.go

CMD ["/build/main"]