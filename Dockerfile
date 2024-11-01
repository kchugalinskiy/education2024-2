FROM golang:1.22-alpine

WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=0 /app/main /app/

ENTRYPOINT ["./main"]