FROM golang as builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn"

# 移动到工作目录：/build
WORKDIR /app

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件app
RUN go build -o shorturl .

# 声明服务端口
EXPOSE 9998

#image stage
FROM scratch
COPY --from=builder /app/go_web /
CMD ["/shorturl"]