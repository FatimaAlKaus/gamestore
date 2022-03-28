FROM golang:1.17-alpine
ENV GIN_MODE=release
WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download

COPY . /app
RUN go build -o app cmd/main.go
ENTRYPOINT ["/app/app"]

