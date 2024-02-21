package config

type Config struct {
	Redis
	Server
	PostgresDb
}

type Redis struct {
	Address  string
	Password string
	Database string
}

type Server struct {
	Port string
}

type PostgresDb struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}
