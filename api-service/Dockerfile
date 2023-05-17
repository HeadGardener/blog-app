FROM golang:1.20-alpine

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o api-service cmd/api-service/main.go

CMD ["./api-service"]