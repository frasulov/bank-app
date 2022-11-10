package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func DBConn() (*sql.DB, error) {
	dbDriver := "postgres"
	dbSource := "postgresql://postgres:postgres@localhost:5433/bank-db?sslmode=disable"
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}
	return testDB, nil

	//host := Configuration.Database.Host
	//dbPort := Configuration.Database.Port
	//user := Configuration.Database.User
	//dbName := Configuration.Database.DBName
	//password := Configuration.Database.Password
	//maxIdleConn := Configuration.Database.GormMaxIdleConn
	//maxOpenConn := Configuration.Database.GormMaxOpenConn
	//maxConnLife := Configuration.Database.GormMaxConnLifetimeHour
	//sslMode := Configuration.Database.SSLMode
	//
	//dbURI := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
	//	user, password, host, dbPort, dbName, sslMode)
	//db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{NamingStrategy: schema.NamingStrategy{
	//	TablePrefix: Configuration.Database.Schema}})
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//sqlDB, err := db.DB()
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//sqlDB.SetMaxIdleConns(maxIdleConn)
	//sqlDB.SetMaxOpenConns(maxOpenConn)
	//sqlDB.SetConnMaxLifetime(time.Duration(maxConnLife))
	//
	//return db, err
}
