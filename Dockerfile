FROM golang:1.21.1-buster as build

WORKDIR /app

COPY go.mod go.sum ./

RUN  go mod download

COPY . .

RUN go build -ldflags='-linkmode external -extldflags -static' -tags netgo -o ./cms ./cmd/server/main.go

FROM scratch 

COPY --from=build /app/cms cms

CMD ["/cms"]

