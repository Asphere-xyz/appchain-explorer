package gateway

import (
	"github.com/Ankr-network/ankr-protocol/shared/common"
	"github.com/spf13/viper"
)

type Config struct {
	GrpcAddress string
	HttpAddress string
}

func (c *Config) ParseFromViper(v *viper.Viper) error {
	c.GrpcAddress = common.ViperGetOrDefault(v, "gateway.grpc-address", ":6565")
	c.HttpAddress = common.ViperGetOrDefault(v, "gateway.http-address", ":8000")
	return nil
}
