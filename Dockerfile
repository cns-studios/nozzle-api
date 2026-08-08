FROM golang:1.26.5-alpine AS builder
RUN apk add --no-cache build-base git ca-certificates
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/nozzle-api .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=builder /app/nozzle-api /app/nozzle-api

ENV PORT=8081
EXPOSE 8081
CMD ["/app/nozzle-api"]