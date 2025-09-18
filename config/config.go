package config

import "github.com/ilyakaznacheev/cleanenv"

type Postgres struct {
	Dsn string `toml:"dsn"`
}

type Config struct {
	Db Postgres `toml:"postgres"`
}

func LoadConfig(file string) (Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(file, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
