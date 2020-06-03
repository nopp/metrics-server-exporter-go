FROM golang:1.13.10 AS builder

ADD go.mod .
ADD main.go .
ADD pod/pod.go ./pod/pod.go
ADD api/api.go ./api/api.go
ADD node/node.go ./node/node.go
ADD config/config.go ./config/config.go

RUN unset GOPATH \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM golang:1.13.10

LABEL maintainer="Carlos Augusto Malucelli <camalucelli@gmail.com>"

COPY --from=builder /go/metrics-server-exporter-go .
ENTRYPOINT ["./metrics-server-exporter-go"]
