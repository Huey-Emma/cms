FROM golang:1.21.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN  go mod download

COPY . .

CMD ["/cms"]

