FROM golang:1.25.0-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -ldflags="-w -s" -o car_rental ./cmd/app

FROM alpine:3.22

RUN adduser -D -h /home/appuser appuser
WORKDIR /home/appuser

COPY --from=builder /app/car_rental ./app

USER appuser

EXPOSE 8080
ENTRYPOINT ["./app"]