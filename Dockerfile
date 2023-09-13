FROM golang:1.21.1-alpine as build

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o cms cmd/server/main.go

FROM scratch

COPY --from=build /app/cms cms

CMD ["/cms"]

