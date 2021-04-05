FROM golang:1.16-alpine AS builder
WORKDIR /server
ENV GO111MODULE=on

# COPY go.mod .
# COPY go.sum .
# RUN go mod download
COPY . .
RUN ls
RUN apk add gcc g++ libc-dev
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-X main.protoVer=$protoVer" -mod=vendor -o main


FROM alpine:latest
ARG BUILD_TIME
ARG SHA1_VER


RUN apk update && \
    apk upgrade && \
    apk add --no-cache curl tzdata && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /server
COPY --from=builder /server/main /server/main

RUN ls
ENV SHA1_VER=${SHA1_VER}
ENV BUILD_TIME=${BUILD_TIME}
# RUN addgroup -g 1000 appuser && \
#     adduser -D -u 1000 -G appuser appuser && \
#     chown -R appuser:appuser /server/main
# USER appuse