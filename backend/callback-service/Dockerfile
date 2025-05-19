FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN apk add --no-cache git && \
    go mod tidy

COPY . ./

RUN go build -o main ./cmd/main.go

EXPOSE 8082

CMD ["./main"]
