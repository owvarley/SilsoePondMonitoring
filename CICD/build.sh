#!/bin/bash -e

# $1 == The Project Path to build

# Get the version for the project to use for the tag
ProjectPath=$1
ProjectVersion=$(cat $ProjectPath/VERSION)
ProjectPathLower=$(echo $ProjectPath | tr '[:upper:]' '[:lower:]')

docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t owenvarley/silsoe$ProjectPathLower:$ProjectVersion --push .
