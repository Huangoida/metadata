#!/bin/bash

rm -rf output
mkdir -p output/conf
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/metadata
cp conf/*.yaml output/conf
tar czvf metadata.tar output
docker buildx build --platform linux/amd64 -t registry.cn-hangzhou.aliyuncs.com/zju_api/metadata:v0.0.1 .
docker push registry.cn-hangzhou.aliyuncs.com/zju_api/metadata:v0.0.1
rm -rf metadata.tar