FROM golang:1.19rc1 AS build

WORKDIR /go/remote-cache

COPY go.mod go.sum ./

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

RUN go env -w GO111MODULE=on \
    && go mod download

COPY ./pkg /go/remote-cache/pkg
COPY ./cmd /go/remote-cache/cmd

RUN go build -o /opt/remote-cache cmd/remote-cache/main.go

FROM library/ubuntu:20.04

COPY --from=build /opt/remote-cache /usr/local/bin/remote-cache


ENTRYPOINT ["/usr/local/bin/remote-cache"]
