FROM golang:1.21-alpine AS builder
WORKDIR /go-app

# install dependencies
COPY ./go.mod ./go.sum /go-app/
RUN go mod download

# move code
COPY ./cmd /go-app/cmd
COPY ./internal /go-app/internal
COPY ./pkg /go-app/pkg

# build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN ls ./cmd && go build -ldflags="-s -w" -o ./bin/main ./cmd \
    && chmod +x ./bin/main
# ---

FROM alpine:3.19 as reservation
COPY --from=builder /go-app/bin/main /main
ENTRYPOINT ["/main", "reservations"]

FROM alpine:3.19 as room
COPY --from=builder /go-app/bin/main /main
ENTRYPOINT ["/main", "rooms"]

FROM alpine:3.19 as stay
COPY --from=builder /go-app/bin/main /main
ENTRYPOINT ["/main", "stays"]

FROM alpine:3.19 as frontdesk
COPY --from=builder /go-app/bin/main /main
ENTRYPOINT ["/main", "frontdesk"]
