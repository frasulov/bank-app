migrationup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable" -verbose up

migrationdown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

startapp:
	go run server.go

swagger:
	swag init -g server.go

mock:
	mockgen --package mockdb --destination db/mock/repository.go BankApp/db/sqlc Repository

mockservice:
	mockgen --package mockdb --destination db/mock/account_service.go BankApp/account AccountService

test:
	go test -v -cover ./...