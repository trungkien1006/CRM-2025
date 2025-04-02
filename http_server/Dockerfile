# Sử dụng image Golang chính thức
FROM golang:1.23-alpine AS builder

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Copy toàn bộ mã nguồn vào container
COPY . .

# Tải các thư viện cần thiết
RUN go mod tidy -v
RUN apk add --no-cache git build-base tzdata

# Biên dịch HTTP server và WebSocket server thành hai binary riêng biệt
RUN go build -o http_server main.go

# Tạo image nhỏ gọn chỉ chứa binary Golang
FROM alpine:latest

# Cài đặt lại tzdata trong runtime image
RUN apk add --no-cache tzdata

# Thiết lập thư mục làm việc
WORKDIR /root/

# Copy binary từ builder
COPY --from=builder /app/http_server .

# Thiết lập timezone mặc định
ENV TZ=Asia/Ho_Chi_Minh

# Mở cổng 8000 cho HTTP server và 8001 cho WebSocket server
EXPOSE 8000

# Chạy cả HTTP server và WebSocket server
CMD ["./http_server"]
