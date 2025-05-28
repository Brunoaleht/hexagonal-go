FROM golang:1.16

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"

RUN go install github.com/spf13/cobra-cli@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0

CMD [ "tail", "-f", "/dev/null" ]
# This Dockerfile sets up a Go development environment with Cobra and Mockgen installed.