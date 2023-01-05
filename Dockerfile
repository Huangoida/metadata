FROM golang:alpine AS builder
WORKDIR /
ADD metadata.tar /opt
WORKDIR /opt/output
CMD ["./metadata"]