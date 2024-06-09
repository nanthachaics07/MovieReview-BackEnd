FROM golang:1.22.2

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

CMD ["air", "cmd/main.go", "-b", "0.0.0.0"]

# Build stage
# FROM golang:1.22.2 AS builder-stage

#   WORKDIR /app
#   ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
#   COPY go.mod go.sum ./
#   RUN go mod download
#   COPY . .
#   RUN go build -o main ./cmd/main.go

# Runner stage
# FROM alpine:latest

#   WORKDIR /root/
#   RUN apk add --no-cache tzdata
#   COPY --from=builder-stage /app/main .
#   EXPOSE 8080

#   CMD ["./main"]
