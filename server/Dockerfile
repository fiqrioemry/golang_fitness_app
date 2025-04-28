FROM golang:1.23.4-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/main.go

RUN chmod +x main

CMD ["go", "run", "./cmd/main.go"]

# CMD ["./main"]
