FROM golang:1.15.0

RUN go get github.com/hpcloud/tail

COPY syslog-tail /opt/syslog-tail

RUN cd /opt/syslog-tail && \
    go build && \
    mv syslog-tail /usr/local/go/bin

WORKDIR /opt