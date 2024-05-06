FROM registry.access.redhat.com/ubi9/ubi-minimal:latest

USER root

RUN microdnf install -y golang-bin

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o rest-api

EXPOSE 8080

CMD ["./rest-api"]
