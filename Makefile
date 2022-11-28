db_host := localhost
db_port := 5432
db_user := root
db_password := secret
db_name := bank

migrationup:
	migrate -path db/migration -database "postgresql://${db_user}:${db_password}@${db_host}:${db_port}/${db_name}?sslmode=disable" -verbose up

migrationup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5435/bank?sslmode=disable" -verbose up 1

migrationdown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5435/bank?sslmode=disable" -verbose down

migrationdown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5435/bank?sslmode=disable" -verbose down 1

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

dockerbuild:
	docker rmi -f bankapp:v1 & docker build -t bankapp:v1 .

dockerrm:
	docker container rm -f bank

dockerrun:
	docker run --name bank --network bank-net -p 8001:8001 bankapp:v1

test:
	go test -v -cover) ./...