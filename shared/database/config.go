package database

import (
	"github.com/spf13/viper"
)

type Config struct {
	PostgresUrl string
}

func (c *Config) ParseFromViper(v *viper.Viper) error {
	c.PostgresUrl = viperGetOrDefault(v, "database.postgres-url", "postgres://postgres:@rpc.dev-01.bas.ankr.com:7432/blockscout?sslmode=disable")
	return nil
}

func viperGetOrDefault(v *viper.Viper, key string, defaultValue string) string {
	if res := v.Get(key); res == nil {
		return defaultValue
	}
	return v.GetString(key)
}
