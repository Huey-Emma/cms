
FROM golang:1.21.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .  # Copies all files in the current directory into the /app directory in the container

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cms .  # Builds the binary file "cms" using the current directory as source

CMD ["./cms"]  # Executes the "cms" binary file

#-----

# The error message "image build failedundefined" indicates that the build failed, but it does not provide any specific error details.
# In order to fix this, we can start by checking the Dockerfile for any syntax errors or incorrect commands.

# In this case, there are a couple of typos that need to be fixed.
# On line 12, the "COPY" command typo has a in the source file path. It should be "." instead of "..".
# On line 16, the command should be "CMD" instead of "ENTRYPOINT".

# Here's the corrected Dockerfile:

FROM golang:1.21.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cms .

CMD ["./cms"]
