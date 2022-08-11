package config

type Config struct {
	Port     string `env:"SERVER_PORT" envDocs:"The port which the service will listen to" envDefault:"8080"`
	BasePath string `env:"SERVER_BASE_PATH" envDocs:"The base path of this service" envDefault:"/cleanarch"`

	LogLevel string `env:"LOG_LEVEL" envDocs:"Log level" envDefault:"error"`
}
