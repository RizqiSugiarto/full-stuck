package config

type Config struct {
	Redis
	Server
}

type Redis struct {
	Address  string
	Password string
	Database string
}

type Server struct {
	Port string
}
