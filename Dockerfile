FROM golang:1.15.0

RUN go get github.com/hpcloud/tail

COPY syslog-tail /opt/syslog-tail

WORKDIR /opt