package test

import (
	"BankApp/config"
	db "BankApp/db/sqlc"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

// install github.com/lib/pq
const (
	dbDriver = "postgres"
)

var testQueries *db.Queries
var testDB *sql.DB
var repository db.Repository

//var accountRepository *account.AccountRepository

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, config.Configuration.Database.URL())
	if err != nil {
		log.Fatal("cannot connect to DB")
	}
	testQueries = db.New(testDB)
	repository = db.NewRepository(testDB)
	os.Exit(m.Run())
}
