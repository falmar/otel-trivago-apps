#!/bin/bash

protoc \
  --proto_path=api/proto/v1 \
  --go_out=./pkg/proto/v1 --go_opt=paths=source_relative \
  --go-grpc_out=./pkg/proto/v1 --go-grpc_opt=paths=source_relative \
  ./api/proto/v1/reservationpb/reservation.proto \
  ./api/proto/v1/roompb/room.proto \
  ./api/proto/v1/staypb/stay.proto
