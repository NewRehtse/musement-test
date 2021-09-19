# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY ./config/* ./config/
COPY ./logger/* ./logger/
COPY ./musement/* ./musement/
COPY ./utils/* ./utils/
COPY ./weather/* ./weather/
COPY ./config.json .

RUN go build -o ./mus

CMD ["./mus"]
