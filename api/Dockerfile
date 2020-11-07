FROM golang:1.15-buster AS builder

COPY . /go/src

WORKDIR /go/src/api

RUN CGO_ENABLED=0 go build 
RUN apt-get install ca-certificates


FROM busybox

WORKDIR /usr/local/bin

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/api/api /api

EXPOSE 50051

CMD ["/api"]
