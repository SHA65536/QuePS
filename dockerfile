# Use the official Golang image to build the Go program
FROM golang:1.22-alpine as builder

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /queps .

FROM alpine:3.18 as run
WORKDIR /app

COPY --from=builder /queps ./queps

CMD ["./queps"]