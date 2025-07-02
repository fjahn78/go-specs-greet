FROM golang:1.24.4-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/grpc_server/*.go

EXPOSE 50051
CMD [ "./svr" ]