package server

type Config struct {
	Host string `toml:"host"`
	Port string `toml:"port"`

	DB_username string `toml:"db_username"`
	DB_password string `toml:"db_password"`
	DB_name     string `toml:"db_name"`
	DB_sslmode  string `toml:"db_sslmode"`
}

func NewConfig() *Config {
	return &Config{
		Host: "127.0.0.1",
		Port: ":8080",

		DB_username: "postgres",
		DB_password: "postgres",
		DB_name: "messanger",
		DB_sslmode: "disable",
	}
}
