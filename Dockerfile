FROM golang:latest

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o cv-builder .

EXPOSE 8000

CMD ["./cv-builder"]