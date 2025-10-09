package config

import "github.com/ilyakaznacheev/cleanenv"

type Postgres struct {
	Dsn string `toml:"dsn"`
}

type Jwt struct {
	Key string `toml:"key"`
}

type Config struct {
	Db  Postgres `toml:"postgres"`
	Jwt Jwt      `toml:"jwt"`
}

func LoadConfig(file string) (Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(file, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
