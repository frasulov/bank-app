package config

import "time"

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
	Host                    string
	Port                    int
	Dialect                 string
	User                    string
	DBName                  string
	Password                string
	GormMaxIdleConn         int
	GormMaxOpenConn         int
	GormMaxConnLifetimeHour int
	SSLMode                 string
	Schema                  string
}
