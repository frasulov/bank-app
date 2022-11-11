
## Simple Bank Application

### Make file
* Above commands are very long. For not wasting time, let's use makefile.
* You need to following credentials in makefile:
  * db connections in migrations commands
  * **-path** in migrations commands
* Commands
  * **migrationup** - instead of ```migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable" -verbose up```
  * **migrationdown** - instead of ```migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable" -verbose down```
  * **startapp** - instead of ```go run server.go```


### Test
* go install github.com/golang/mock/mockgen@v1.6.0
* mockgen --package mockdb --destination db/mock/repository.go BankApp/db/sqlc Repository

### References
* Validations - https://github.com/go-playground/validator
* Swagger - https://github.com/swaggo/swag
* Fiber - https://docs.gofiber.io
* Migration - https://github.com/golang-migrate/migrate
* Sqlc - https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html