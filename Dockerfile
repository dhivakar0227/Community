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

RUN go get go.mongodb.org/mongo-driver/mongo

RUN go env

ADD . /go/src/github.com/dhivakar0227/Community

RUN ls /go/src/github.com/dhivakar0227/Community

RUN go install github.com/dhivakar0227/Community/src/servers/Questions

RUN ls /go/bin/Questions

ENTRYPOINT ["/go/bin/servers"]

EXPOSE 8081
