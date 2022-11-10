
## Default Project Structure with SQLC library

### Create new app

* Copy 'sample' directory for your each model
* Rename 'sample' word with your model name
* Add your service functions in your-model/service.go
* Add your built in repository function in query/your-model.sql & run sqlc generate
* Add your custom repository functions in your-model/repository.go
* Add dto (with validations) in your-model/dto.go

### Required Libraries & Actions
* github.com/lib/pq library must be imported in all files that db connection created
```go get github.com/lib/pq```
* run following command for migrations
  * db/migration is folder that your migration files will be stored. init_schema word in command depends on you. It will be added to migration file name after its version, and generally it means what your sql file contains. Command:
  ```migrate create -ext sql -dir db/migration -seq init_schema ```
  * Run following command to run the latest up migration file.
  ```migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable" -verbose up```
  * Run following command to run the latest up migration file.
  ```migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable" -verbose down```
* Install sqlc:
  ```https://docs.sqlc.dev/en/latest/overview/install.html```
* run following command if there is no sqlc.yaml in your main project folder. 
```sqlc init```
* create query in the folder that you specifies in sqlc.yaml: ```queries: "./db/query/"```. In our case, I will create queries in query folder within db folder.
* after writing your queries run following command:
```sqlc generate```
* Now, your repository functions created, you can use it in your service.
* Moreover, you can create custom repository function in sample/repository.go

### Make file
* Above commands are very long. For not wasting time, let's use makefile.
* You need to following credentials in makefile:
  * db connections in migrations commands
  * **-path** in migrations commands
* Commands
  * **migrationup** - instead of ```migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable" -verbose up```
  * **migrationdown** - instead of ```migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable" -verbose down```
  * **startapp** - instead of ```go run server.go```

### In Service
* cast model to dto (using mappers)
* return dto
* return 'HttpError' as error. Following example:
```
import my_errors "default-project-structure/errors"
my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("not_found", "en"))
```
* 'HttpError' support multi language also, you can add languages in errors/errorResponse.json file

### In Route
* use Routing constraints ([https://docs.gofiber.io/guide/routing](https://docs.gofiber.io/guide/routing))


### In Middleware
* add all common properties in Common function
* add your protect endpoint logic in Protect function
* Example adding Protect middleware
```
- app.Get("/samples/:id<int>", controller.GET)
+ app.Get("/samples/:id<int>", middleware.Protect(), controller.GET)
```

### In Controller
* fill POST, GET, PUT, PATCH, DELETE function whenever needed
* you can add also your custom methods as function to your controller


### Swagger
* swagger installed
* run following command after each swagger comment added:
```
swag init -g server.go
```


### Transactions
* It is used to write transactions. The example functions already added. How to use it in your service? Following steps
* Uncomment store in service struct
* Uncomment store in GetNewSampleService function
* Add 'store *db.Store' param to GetNewSampleService function
* Uncomment line 38 ```store := db.NewStore(connection)``` in server.go
* Add store as param to following function ```sampleService := sample.GetNewSampleService(sampleRepository, store)```

### References
* Validations - https://github.com/go-playground/validator
* Swagger - https://github.com/swaggo/swag
* Fiber - https://docs.gofiber.io
* Migration - https://github.com/golang-migrate/migrate
* Sqlc - https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html