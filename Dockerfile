
FROM golang:1.21.1-alpine

WORKDIR /app

COPY go.mod go.sum ./  # Corrected typo in the source file path

RUN go mod download

COPY . .  # Copies all files in the current directory into the /app directory in the container

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cms .  # Builds the binary file "cms" using the current directory as source

CMD ["./cms"]  # Executes the "cms" binary file
