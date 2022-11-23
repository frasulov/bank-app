# build stages
FROM golang:1.18-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main-bank-app server.go

# run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main-bank-app .
EXPOSE 8001
CMD [ "/app/main" ]