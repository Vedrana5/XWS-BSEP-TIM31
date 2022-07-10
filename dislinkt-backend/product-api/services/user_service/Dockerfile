FROM golang:latest as builder
WORKDIR /app

COPY ./user_service/go.mod ./
COPY ./user_service/go.sum ./
COPY /common ../common
RUN go mod download
COPY ./user_service/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /root/


COPY --from=builder /app/main .

EXPOSE 9091
CMD ["./main"]