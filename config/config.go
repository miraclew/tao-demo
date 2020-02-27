package config

import (
	"errors"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Env       string `env:"ENV" envDefault:"prod"`
	HttpBind  string `env:"HTTP_BIND" envDefault:":8080"`
	RedisAddr string `env:"REDIS"`
	MysqlAddr string `env:"MYSQL"`

	JwtAlgorithm string `env:"JWT_ALGORITHM" envDefault:"HS256"`
	JwtSecret    string `env:"JWT_SECRET"`
	JwtExpiresIn int64  `env:"JWT_EXPIRES_IN" envDefault:"3000"`


	LogLevel  string `env:"LOG_LEVEL" envDefault:"info"`
	LogFormat string `env:"LOG_FORMAT" envDefault:"text"`
}

func (o Config) Validate() error {
	if o.RedisAddr == "" {
		return errors.New("require redis address")
	}
	if o.MysqlAddr == "" {
		return errors.New("require mysql address")
	}
	if o.JwtSecret == "" {
		return errors.New("require jwt secret")
	}

	return nil
}

func NewConfig() *Config {
	// load .env into environment
	configFile := ".env"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	if _, err := os.Stat(configFile); err == nil {
		if err := godotenv.Load(configFile); err != nil {
			panic(err)
		}
	}

	conf := new(Config)
	err := env.Parse(conf)
	if err != nil {
		panic(err)
	}
	return conf
}
