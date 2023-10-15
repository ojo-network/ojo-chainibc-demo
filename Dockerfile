# Builder
FROM golang:1.20-alpine AS builder

RUN apk add --no-cache \
    ca-certificates \
    build-base \
    linux-headers \
    git

WORKDIR /chainibcdemo

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

# Runner
FROM alpine
RUN apk add bash

COPY --from=builder /chainibcdemo/build/chainibcdemo /bin/chainibcdemo

EXPOSE 26656 26657 1317 9090
ENTRYPOINT ["chainibcdemo", "start"]
