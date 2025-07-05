FROM golang:1.24

# Downloading dependencies and only redownloading them in subsequent builds if they change
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copying source files and compiling them
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./src/api/main.go

# Running binary
CMD ["/main"]