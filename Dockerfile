FROM golang:1.11 as builder

## Create a directory and Add Code
RUN mkdir -p /go/src/github.com/orvice/squid-log-server
WORKDIR /go/src/github.com/orvice/squid-log-server
ADD .  /go/src/github.com/orvice/squid-log-server

RUN go get
RUN go build -o squid-log-server


FROM orvice/go-runtime

COPY --from=builder /go/src/github.com/orvice/squid-log-server/squid-log-server .

EXPOSE 8801 8802
ENTRYPOINT [ "./squid-log-server" ]