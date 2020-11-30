FROM golang:1.15-alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# 将代码复制到容器中
COPY . .

RUN go build -o app .

FROM scratch

COPY --from=builder /build/app /

ENTRYPOINT ["/app"]