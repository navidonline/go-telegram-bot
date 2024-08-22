FROM golang:bullseye AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server

CMD [ "/app/server" ]