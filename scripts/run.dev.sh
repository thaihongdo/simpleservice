#!/bin/sh

TAG="latest"
IMAGE="simpleservice"
echo "~~~~~ Starting build developement entity-server ${IMAGE} ${TAG} ~~~~~~~~~"
docker build -t $IMAGE:$TAG  -f ./build/.docker/dev.dockerfile .
echo '~~~~~ Finish build developement entity-server ~~~~~~~~~'
docker run -d -p 8080:8080 $IMAGE:$TAG
echo '~~~~~ Finish run developement entity-server ~~~~~~~~~'