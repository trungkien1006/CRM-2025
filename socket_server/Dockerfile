FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod và go.sum trước để tận dụng cache Docker
COPY go.mod go.sum ./

# Tải dependencies trước
RUN go mod download

# Copy toàn bộ source code vào container
COPY . .

# Kiểm tra và dọn dẹp dependencies
RUN go mod tidy

# Biên dịch ứng dụng
RUN go build -o websocket_server

# Tạo một image tối ưu để chạy ứng dụng
FROM alpine:latest

WORKDIR /root/

# Cài đặt thư viện cần thiết (nếu có)
RUN apk add --no-cache ca-certificates

# Copy binary từ builder
COPY --from=builder /app/websocket_server .

# Copy file .env nếu cần
# COPY .env .

# Expose port của WebSocket server
EXPOSE 8001

# Chạy ứng dụng
CMD ["./websocket_server"]
