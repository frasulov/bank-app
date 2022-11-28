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
	//dbSource1 := "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable"
	dbSource := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.SSLMode)
	fmt.Println("db: ", dbSource)
	testDB, err := sql.Open(dbDriver, dbSource)
	testDB.SetMaxIdleConns(dbConfig.GormMaxIdleConn)
	testDB.SetMaxOpenConns(dbConfig.GormMaxOpenConn)
	testDB.SetConnMaxLifetime(time.Duration(dbConfig.GormMaxConnLifetimeHour))
	if err != nil {
		fmt.Println("err happens: ", err)
		return nil, err
	}
	return testDB, nil
}
