# install golang
FROM golang

# install protobuf from source
RUN apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool

# NOTE: for now, this docker image always builds the current HEAD version of
# gRPC.  After gRPC's beta release, the Dockerfile versions will be updated to
# build a specific version.

# Get the source from GitHub
RUN go get google.golang.org/grpc
# Install protoc-gen-go
RUN go get github.com/golang/protobuf/protoc-gen-go

RUN go env

ADD . /go/src/github

RUN ls /go/src/github/Community

RUN go install github/Community/src/Servers/Questions

RUN ls /go/bin/Servers

ENTRYPOINT ["/go/bin/Servers"]

EXPOSE 8081