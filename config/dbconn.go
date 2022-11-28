package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func DBConn() (*sql.DB, error) {
	dbDriver := "postgres"
	dbConfig := Configuration.Database
	testDB, err := sql.Open(dbDriver, dbConfig.URL())
	testDB.SetMaxIdleConns(dbConfig.MaxIdleConn)
	testDB.SetMaxOpenConns(dbConfig.MaxOpenConn)
	testDB.SetConnMaxLifetime(time.Duration(dbConfig.MaxConnLifetimeHour))
	if err != nil {
		fmt.Println("err happens: ", err)
		return nil, err
	}
	return testDB, nil
}
