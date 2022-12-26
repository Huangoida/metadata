#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build -o output/metadata
docker buildx build --platform linux/amd64 -t registry.cn-hangzhou.aliyuncs.com/zju_api/metadata:v0.0.1 .
docker push registry.cn-hangzhou.aliyuncs.com/zju_api/metadata:v0.0.1