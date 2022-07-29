FROM golang:alpine as dev

WORKDIR /app

COPY . .

# install protoc for grpc
RUN apk add protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN go mod tidy

CMD [ "go","run","./cmd" ]

# =============== #

FROM dev as builder

WORKDIR /app

RUN go build -o ./build/main ./cmd/main.go

# =============== #

FROM alpine:latest as prod 

WORKDIR /app

COPY --from=builder ./build/main ./

CMD [ "./main" ]
