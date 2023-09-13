
FROM golang:1.21.1-buster as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cms ./cmd/server/main.go

FROM alpine:latest

COPY --from=build /app/cms cms

CMD ["/cms"]
