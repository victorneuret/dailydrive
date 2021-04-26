# Build stage
FROM golang:1.16.2-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o binary .

# Run stage
FROM scratch

WORKDIR /app

COPY --from=builder /app/binary ./
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /etc/ssl/private /etc/ssl/private

ENTRYPOINT ["./binary"]

