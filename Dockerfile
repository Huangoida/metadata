FROM golang:alpine AS builder
WORKDIR /
COPY output/metadata /opt
WORKDIR /opt
CMD ["./metadata"]