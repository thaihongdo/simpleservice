ARG GO_VERSION=1.18
FROM golang:${GO_VERSION}-alpine AS build_base
LABEL stage=build_base
RUN apk update && apk add gcc libc-dev make git --no-cache ca-certificates  && \
    mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN mkdir -p /go/src/simpleservice
WORKDIR /go/src/simpleservice

ENV GO111MODULE=on
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# ================ copy from stage build ===========

FROM build_base AS server_builder
LABEL stage=server_builder

RUN mkdir -p /go/src/simpleservice/cmd
COPY ./cmd/entity-server /go/src/simpleservice/cmd/entity-server

RUN mkdir -p /go/src/simpleservice/internal
COPY ./internal/pkg /go/src/simpleservice/internal/pkg

RUN mkdir -p /go/src/simpleservice/configs
COPY ./configs /go/src/simpleservice/configs

RUN mkdir -p /go/src/simpleservice/storage
COPY ./storage /go/src/simpleservice/storage

WORKDIR /go/src/simpleservice/cmd/entity-server/

RUN go build  -o /entity-server .

# ================ copy from stage build ===========
FROM alpine:3.8

RUN apk update &&  apk add --no-cache ca-certificates git && \
    mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN mkdir -p /simpleservice
WORKDIR /simpleservice

RUN mkdir -p /simpleservice/cmd/bin
RUN mkdir -p /simpleservice/configs
RUN mkdir -p /simpleservice/storage

COPY --from=server_builder /entity-server /simpleservice/cmd/bin/
COPY --from=server_builder /go/src/simpleservice/configs/config.toml /simpleservice/configs/
COPY --from=server_builder /go/src/simpleservice/storage /simpleservice/storage/

RUN chmod -R 777 /simpleservice/cmd/bin

RUN chown -R nobody:nobody /simpleservice
RUN chmod -R 755 /simpleservice

USER nobody:nobody

EXPOSE 8080
ENTRYPOINT ["/simpleservice/cmd/bin/entity-server"]