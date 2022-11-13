migrationup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable" -verbose up

migrationup1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable" -verbose up 1

migrationdown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable" -verbose down

migrationdown1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable" -verbose down 1

migrationnew:
	migrate create -ext sql -dir db/migration -seq add_users

sqlc:
	sqlc generate

server:
	go run server.go

swagger:
	swag init -g server.go

mock:
	mockgen --package mockdb --destination db/mock/account_service.go BankApp/account AccountService

test:
	go test -v -cover ./...