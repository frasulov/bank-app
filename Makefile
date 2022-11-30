DB_HOST := localhost
DB_PORT := 5435
DB_USER := root
DB_PASSWORD := secret
DB_NAME := bank
DB_SSL_MODE := disable
DB_SOURCE := "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}"

migrationup:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose up

migrationup1:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose up 1

migrationdown:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose down

migrationdown1:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose down 1

migrationnew:
	migrate create -ext sql -dir db/migration -seq add_sessions

sqlc:
	sqlc generate

server:
	go run server.go

swagger:
	swag init -g server.go

mock:
	mockgen --package mockdb --destination db/mock/account_service.go BankApp/account AccountService

dockerbuild:
	docker rmi -f bankapp:v1 & docker build -t bankapp:v1 .

dockerrm:
	docker container rm -f bank

dockerrun:
	docker run --name bank --network bank-net -p 8001:8001 bankapp:v1

test:
	go test -v -cover ./...