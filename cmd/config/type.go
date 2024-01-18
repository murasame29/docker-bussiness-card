package config

type config struct {
	Server struct {
		Port string `env:"SERVER_PORT" envDefault:"8080"`
		Host string `env:"SERVER_HOST" envDefault:"localhost"`
	}
}

var Config *config
