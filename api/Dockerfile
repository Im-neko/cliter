FROM golang:1.15-buster AS builder

COPY . /go/src

WORKDIR /go/src/api

#RUN CGO_ENABLED=0 go build 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -installsuffix "static" .

RUN wget -O /bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.3.1/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

RUN apt-get install ca-certificates

RUN ls
RUN ls /go/bin

FROM busybox

WORKDIR /usr/local/bin

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/api /bin/api
COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe

CMD ["/bin/api"]
