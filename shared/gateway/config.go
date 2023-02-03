package gateway

import (
	"github.com/Ankr-network/appchain-explorer/shared/common"
	"github.com/spf13/viper"
)

type Config struct {
	GrpcAddress  string
	HttpAddress  string
	StaticFolder string
}

func (c *Config) ParseFromViper(v *viper.Viper) error {
	c.GrpcAddress = common.ViperGetOrDefault(v, "gateway.grpc-address", ":6565")
	c.HttpAddress = common.ViperGetOrDefault(v, "gateway.http-address", ":8000")
	c.StaticFolder = common.ViperGetOrDefault(v, "gateway.static-folder", "./web/build")
	return nil
}
