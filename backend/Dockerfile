FROM golang:latest
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ .
RUN go build -o main ./cmd/main.go

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]