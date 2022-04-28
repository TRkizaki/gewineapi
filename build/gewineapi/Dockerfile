
   
FROM golang:1.17.1-alpine as builder

WORKDIR /build
COPY ../../go.mod ../../go.sum ./
RUN go mod download
COPY ../../  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -ldflags '-s -w' ./cmd/gewineapi

FROM alpine
COPY --from=builder /build/gewineapi /opt/app/
ENTRYPOINT ["/opt/app/gewineapi"]