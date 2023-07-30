FROM golang:1.19-alpine AS builder
WORKDIR /go-app

# install dependencies
COPY ./go.mod ./go.sum /go-app/
RUN go mod download

# move code
COPY ./cmd /go-app/cmd
COPY ./internal /go-app/internal
COPY ./pkg /go-app/pkg

# build
# TODO: move each go build to its own stage for parallelism
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN ls ./cmd && go build -ldflags="-s -w" -o ./bin/reservation ./cmd/reservationsvc \
    && chmod +x ./bin/reservation \
    && go build -ldflags="-s -w" -o ./bin/room ./cmd/roomsvc \
    && chmod +x ./bin/room
# ---

FROM alpine:3.18 as reservation
COPY --from=builder /go-app/bin/reservation /main
ENTRYPOINT ["/main"]


FROM alpine:3.18 as room
COPY --from=builder /go-app/bin/room /main
ENTRYPOINT ["/main"]
