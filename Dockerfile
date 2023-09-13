
FROM golang:1.21.1-buster as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd/server/main.go .  # Fixed typo in the source file path

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cms main.go  # Removed unnecessary directory path

FROM alpine:latest

COPY --from=build /app/cms cms

CMD ["/cms"]
