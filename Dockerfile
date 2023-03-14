FROM golang:1.20-alpine

WORKDIR /app

ENV TOKEN = ""

COPY . .

RUN go build .

CMD ./leimdorfer-bot $TOKEN
