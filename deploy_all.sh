#!/bin/bash
set -e

echo "🔧 Building HTTP server..."
cd /home/doanvanquoc/CRM-2025/http_server
/usr/local/go/bin/go build -o http_server main.go
chmod +x http_server

echo "🔧 Building Socket server..."
cd ..
cd socket_server
/usr/local/go/bin/go build -o socket_server socket.go
chmod +x socket_server

echo "🔧 Building Kafka server..."
cd ..
cd kafka_server
/usr/local/go/bin/go build -o kafka_server consumer.go
chmod +x kafka_server

echo "🛑 Stopping all services..."
sudo systemctl stop kafka_server
sudo systemctl stop socket_server
sudo systemctl stop http_server

echo "📥 Pulling latest code..."
cd ..
git pull origin main

echo "✅ Starting services..."
sudo systemctl start kafka_server
sudo systemctl start socket_server
sudo systemctl start http_server

echo "🎉 Triển khai hoàn tất!"
