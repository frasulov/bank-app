package config

type ProfileConfigurations struct {
	Profile Profile
}

type Profile struct {
	Active string
}

type Configurations struct {
	App
	Database
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
