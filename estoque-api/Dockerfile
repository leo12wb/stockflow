FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY . .

# instala swag e gera os docs ANTES do go mod tidy (o pacote docs precisa existir)
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init -g cmd/server/main.go --output docs

RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:3.20
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./server"]
