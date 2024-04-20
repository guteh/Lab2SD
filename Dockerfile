FROM golang:latest

WORKDIR /app

COPY ./central /app

RUN go build -o central central.go

EXPOSE 8080

CMD ["./central"]
