FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -o node ./cmd/node
RUN go build -o controller ./cmd/controller

CMD ["./node"]