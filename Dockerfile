FROM golang:1.13.9 AS builder
ENV msg_version 1.0
#RUN curl -LO https://github.com/nopp/metrics-server-exporter-go/archive/${msg_version}.tar.gz \
#    && tar zxf ${msg_version}.tar.gz
#WORKDIR metrics-server-exporter-go-${msg_version}
ADD main.go .
ADD go.mod .
ADD common/common.go ./common/common.go
ADD pod/pod.go ./pod/pod.go
ADD node/node.go ./node/node.go
RUN unset GOPATH \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM golang:1.13.9

LABEL maintainer="Carlos Augusto Malucelli <camalucelli@gmail.com>"

ENV msg_version 1.0
#COPY --from=builder /go/metrics-server-exporter-go-${msg_version}/main .
COPY --from=builder /go/metrics-server-exporter-go .
ENTRYPOINT ["./metrics-server-exporter-go"]
