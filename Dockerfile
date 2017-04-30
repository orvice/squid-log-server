FROM golang:1.8

## Create a directory and Add Code
RUN mkdir -p /go/src/github.com/orvice/squid-log-server
WORKDIR /go/src/github.com/orvice/squid-log-server
ADD .  /go/src/github.com/orvice/squid-log-server

# Download and install any required third party dependencies into the container.
RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8801

# Now tell Docker what command to run when the container starts
CMD ["go-wrapper", "run"]