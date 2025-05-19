FROM golang:1.23-alpine

WORKDIR /app
COPY . .

RUN apk add --no-cache git && \
    go mod tidy && \
    go build -o main ./cmd/main.go

EXPOSE 8086
CMD ["./main"]
