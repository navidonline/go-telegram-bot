FROM golang:bullseye AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=1 GOOS=linux go build -v -o server -a -ldflags '-linkmode external -extldflags "-static"' .

CMD [ "/app/server" ]