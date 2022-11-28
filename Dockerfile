# build stages
FROM golang:1.18-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main-bank-app server.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main-bank-app .
COPY --from=builder /app/migrate ./migrate
COPY config/config-profile.yml /app/config/
COPY config/app.env /app/config/
COPY config/config-dev.yml /app/config/
COPY config/config-prod.yml /app/config/
COPY errors/errorResponse.json /app/errors/
COPY db/migration ./migration
COPY start.sh /app/

EXPOSE 8001
CMD [ "/app/main-bank-app" ]
ENTRYPOINT [ "/app/start.sh" ]