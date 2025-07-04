FROM golang:1.24-alpine AS builder

WORKDIR /app

ARG bin_to_build

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/$bin_to_build/main.go

FROM alpine:latest
COPY --from=builder /app/svr .
CMD [ "./svr" ]