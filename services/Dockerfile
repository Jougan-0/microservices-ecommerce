FROM golang:1.21 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN ls -la /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backend ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/backend .

ENV BUILD_TYPE=dockerfile

EXPOSE 8080

CMD ["./backend"]
