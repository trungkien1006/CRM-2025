FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy toàn bộ source code vào container
COPY . .

# Kiểm tra và dọn dẹp dependencies
RUN go mod tidy -v

# Biên dịch ứng dụng
RUN go build -o kafka_consumer

# Tạo image chạy ứng dụng
FROM alpine:latest

WORKDIR /root/

# Cài đặt thư viện cần thiết
RUN apk add --no-cache ca-certificates

# Copy binary từ builder
COPY --from=builder /app/kafka_consumer .

EXPOSE 8002

CMD ["./kafka_consumer"]
