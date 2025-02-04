FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -o openapi .

EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]
