FROM golang:1.16

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


RUN mkdir go-grpc
WORKDIR /go-grpc

COPY . . 

RUN go build server/main.go

EXPOSE 50051

CMD ./main
