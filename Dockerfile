FROM golang:1.21-alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://goproxy.cn,direct
WORKDIR /build
COPY . .
RUN apk add  git && COMMIT_SHA1=$(git rev-parse --short HEAD || echo "0.0.0") \
    BUILD_TIME=$(date "+%F %T") \
    go build -ldflags="-s -w -X 'github.com/hhstu/gin-template/utils.Version=$1' -X 'github.com/hhstu/gin-template/utils.Commit=${COMMIT_SHA1}' -X 'github.com/hhstu/gin-template/utils.Date=${BUILD_TIME}'"   -o  gin-template cmd/app.go

FROM ubuntu:22.04
COPY --from=builder  /build/gin-template /gin-template
CMD /gin-template