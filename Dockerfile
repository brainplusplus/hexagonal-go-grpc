FROM --platform=linux/arm64 golang:1.24.1-alpine3.21 AS builder
WORKDIR /app
COPY . ./
RUN go mod tidy
RUN GOOS=linux GOARCH=arm64 go build -o app ./cmd/app/main.go

FROM alpine:latest
# Set the working directory inside the container
RUN apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime \
    && echo "Asia/Jakarta" > /etc/timezone
ENV TZ=Asia/Jakarta
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 3000
ENTRYPOINT ["./app"]