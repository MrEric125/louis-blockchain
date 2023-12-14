FROM golang:1.21.4

WORKDIR /app

COPY . /app

# 设置环境变量
ENV CGO_ENABLED=0

RUN go build -o myapp

EXPOSE 8080

CMD ["./myapp"]