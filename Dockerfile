
FROM golang:1.21.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .  # Fixed typo in the source file path

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cms .  # Removed unnecessary directory path

CMD ["./cms"]  # Fixed typo in the command
