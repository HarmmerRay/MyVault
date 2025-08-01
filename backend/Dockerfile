FROM golang:1.21-alpine AS builder

WORKDIR /app

# 安装git和ca-certificates
RUN apk add --no-cache git ca-certificates

# 复制go mod文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# 运行时镜像
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从builder镜像复制二进制文件
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 8081

# 运行应用
CMD ["./main"]