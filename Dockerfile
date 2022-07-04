FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/Makson4iik/Chat_Api/
WORKDIR /go/src/github.com/Makson4iik/Chat_Api
RUN go mod download
COPY . /go/src/github.com/Makson4iik/Chat_Api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/chat_api github.com/Makson4iik/Chat_Api

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/Makson4iik/Chat_Api/build/chat_api /usr/bin/chat_api
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/chat_api"]
