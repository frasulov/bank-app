package config

import (
	"fmt"
	"time"
)

type ProfileConfigurations struct {
	Profile Profile
}

type Profile struct {
	Active string
}

type Configurations struct {
	App
	Database
	Token
}

type Token struct {
	SymmetricKey            string        `mapstructure:"symmetric_key"`
	AccessTokenDuration     time.Duration `mapstructure:"access_token_duration"`
	AuthorizationHeaderKey  string        `mapstructure:"authorization_header_key"`
	AuthorizationTypeBearer string        `mapstructure:"authorization_type_bearer"`
	AuthorizationPayloadKey string        `mapstructure:"authorization_payload_key"`
}

type App struct {
	Host string
	Port int
}

type Database struct {
	Host                string `mapstructure:"DB_HOST"`
	Port                int    `mapstructure:"DB_PORT"`
	User                string `mapstructure:"DB_USER"`
	DBName              string `mapstructure:"DB_NAME"`
	Password            string `mapstructure:"DB_PASSWORD"`
	MaxIdleConn         int    `mapstructure:"DB_MAX_IDLE_CONNECTION"`
	MaxOpenConn         int    `mapstructure:"DB_MAX_OPEN_CONNECTION"`
	MaxConnLifetimeHour int    `mapstructure:"DB_MAX_CONNECTION_LIFETIME"`
	SSLMode             string `mapstructure:"DB_SSL_MODE"`
	Schema              string `mapstructure:"DB_SCHEMA"`
}

func (dbConfig *Database) URL() string {
	dbSource := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.SSLMode)
	return dbSource
}
