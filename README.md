
## Simple Bank Application
* Create Account
* Get Account
* Transfer Money
* ...

### Test
* ```go install github.com/golang/mock/mockgen@v1.6.0```
* ```go get github.com/golang/mock/mockgen@v1.6.0```
* ```mockgen --package mockdb --destination db/mock/repository.go BankApp/db/sqlc Repository```

### References
* Validations - https://github.com/go-playground/validator
* Swagger - https://github.com/swaggo/swag
* Fiber - https://docs.gofiber.io
* Migration - https://github.com/golang-migrate/migrate
* Sqlc - https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html