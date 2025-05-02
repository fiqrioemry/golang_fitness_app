FROM golang:1.23.4-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/main.go

# RUN chmod +x main

# Salin script
COPY wait-for-it.sh /app/wait-for-it.sh

RUN chmod +x /app/wait-for-it.sh

ENTRYPOINT ["/app/wait-for-it.sh", "db:3306", "--"]

CMD ["go", "run", "./cmd/main.go"]

# CMD ["./main"]
