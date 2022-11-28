# build stages
FROM golang:1.18-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main-bank-app server.go

# run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main-bank-app .
COPY config/config-profile.yml /app/config/
COPY config/config-dev.yml /app/config/
COPY config/config-prod.yml /app/config/
COPY errors/errorResponse.json /app/errors/

EXPOSE 8001
CMD [ "/app/main-bank-app" ]