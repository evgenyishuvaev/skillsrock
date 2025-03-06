FROM golang:1.23.6-alpine3.21 AS builder
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./app/main.go
CMD ["/app/main"]
