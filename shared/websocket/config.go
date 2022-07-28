package websocket

import (
	"fmt"
	"github.com/Ankr-network/ankr-protocol/shared/common"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	BlockscoutWebSocketUrl string
	EthereumWebSocketUrl   string
}

func (c *Config) ParseFromViper(v *viper.Viper) error {
	c.BlockscoutWebSocketUrl = common.ViperGetOrDefault(v, "websocket.blockscout-websocket-url", "wss://blockscout.dev-02.bas.ankr.com/socket/websocket?locale=en&vsn=2.0.0")
	c.EthereumWebSocketUrl = common.ViperGetOrDefault(v, "websocket.ethereum-websocket-url", "wss://mainnet.infura.io/ws/v3/261e53a44e3a49eba19eaa54e4907912")
	if !strings.Contains(c.BlockscoutWebSocketUrl, "vsn=") || !strings.Contains(c.BlockscoutWebSocketUrl, "locale=") {
		return fmt.Errorf("there is no (vsn or locale) query, connection might not work")
	}
	return nil
}
