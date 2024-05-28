FROM golang:1.22 as builder

WORKDIR /email-service

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o email-service /email-service/main.go

FROM alpine:latest

WORKDIR /email-service

COPY --from=builder email-service/ /email-service/

EXPOSE 3000

CMD EXECUTION_MODE=indexer ./email-service && ./email-service
