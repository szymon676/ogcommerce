FROM golang:1.20.1-alpine

WORKDIR /app

COPY . .
COPY .env .

RUN go build -o main main.go

EXPOSE 4000

CMD ["./main"]