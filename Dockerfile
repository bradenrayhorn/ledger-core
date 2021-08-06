FROM golang:1.16.4 as build

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o ./ledger-core ./cmd/core

FROM alpine:latest
COPY --from=build /app/ledger-core /app/

EXPOSE 8080

ENTRYPOINT ["/app/ledger-core"]
