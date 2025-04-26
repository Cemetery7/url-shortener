package config

type Config struct {
	Env string `aml:"env" env:"ENV envDefaul: "development"`
}
