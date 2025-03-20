package config

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type JWTConfig struct {
	Secret string
}

func Load() (*Config, error) {
	return &Config{
		Server: ServerConfig{
			Port: 8080,
		},
		Database: DatabaseConfig{
			Driver:   "mysql",
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "123456",
			Database: "dididaren",
		},
		JWT: JWTConfig{
			Secret: "your-secret-key",
		},
	}, nil
}
