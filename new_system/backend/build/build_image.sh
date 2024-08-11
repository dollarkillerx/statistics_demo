#!/bin/sh
image_name=followme
version=1.0.0

docker rmi -f harbor.dollarkiller.com/library/$image_name:$version
docker build -f build/Dockerfile -t harbor.dollarkiller.com/library/$image_name:$version  .
docker push harbor.dollarkiller.com/library/$image_name:$version
#docker save -o $image_name-$version.tar $image_name:$version