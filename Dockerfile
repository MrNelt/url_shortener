FROM golang:1.21-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN mkdir -p /usr/local/bin/
RUN go mod tidy
RUN go build -o /usr/local/bin/app ./cmd/app

EXPOSE 8080

CMD ["/usr/local/bin/app"]