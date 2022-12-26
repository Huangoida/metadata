FROM golang:alpine AS builder
WORKDIR /
ADD metadata.tar /opt
CMD ["/opt/output/metadata"]