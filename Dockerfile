FROM registry.access.redhat.com/ubi9/go-toolset:1.21

USER root

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o rest-api

EXPOSE 8080

CMD ["./rest-api"]
