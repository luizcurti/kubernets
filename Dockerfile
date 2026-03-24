FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server .

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8000
CMD ["./server"]