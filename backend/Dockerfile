FROM golang:latest AS builder
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ .
RUN go build -o main ./cmd/main.go

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/main .
COPY backend/wait-for-it.sh .
RUN chmod +x wait-for-it.sh

CMD ["./wait-for-it.sh", "postgres:5432", "--", "./main"]
