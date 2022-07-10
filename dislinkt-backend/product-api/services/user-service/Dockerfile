FROM golang:latest as builder
WORKDIR /app

COPY ./user-service/go.mod ./
COPY ./user-service/go.sum ./
COPY /common ../common
RUN go mod download
COPY ./user-service/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-service

FROM alpine:latest

WORKDIR /root/


COPY --from=builder /app/user-service .

EXPOSE 9091
CMD ["./user-service"]