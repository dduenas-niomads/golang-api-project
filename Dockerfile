FROM golang:1.22-alpine

WORKDIR /app

COPY /go-app/* ./

RUN go mod download

COPY . .

RUN go run .

EXPOSE 8888

CMD ["./main"]