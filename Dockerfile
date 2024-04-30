FROM golang:1.20.3-alpine AS builder

COPY . /github.com/vadskev/chat-auth/sourse/
WORKDIR /github.com/vadskev/chat-auth/sourse/

RUN go mod download
RUN go build -o ./bin/chat_auth cmd/grpc_server/main.go

FROM alpine:3.14

WORKDIR /root/
COPY --from=builder /github.com/vadskev/chat-auth/sourse/bin/chat_auth .

CMD ["./chat_auth"]