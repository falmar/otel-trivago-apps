#!/bin/bash

LOAD=true
PUSH=$1
if [ -z "$PUSH" ]; then
  PUSH=false
else
  LOAD=false
fi

docker buildx build --load=$LOAD --push=$PUSH -f ./build/Dockerfile \
  -t falmar/otel-trivago:reservation \
  --target=reservation .

docker buildx build --load=$LOAD --push=$PUSH -f ./build/Dockerfile \
  -t falmar/otel-trivago:room \
  --target=room .
