#!/bin/bash

LOAD=true
PUSH=$1
if [ -z "$PUSH" ]; then
  PUSH=false
else
  LOAD=false
fi

images="reservation room stay"

for i in $images; do
  docker buildx build --load=$LOAD --push=$PUSH -f ./build/Dockerfile \
    -t falmar/otel-trivago:$i \
    --target=$i \
    .
done
