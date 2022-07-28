package websocket

import (
	"github.com/Ankr-network/ankr-protocol/shared/common"
	"github.com/spf13/viper"
	"net/url"
)

type Config struct {
	BlockscoutWebSocketUrl string
	EthereumWebSocketUrl   string
}

func fixBlockscoutWebSocketUrl(value string) (string, error) {
	u, err := url.Parse(value)
	if err != nil {
		return "", err
	}
	query := u.Query()
	if !query.Has("locale") {
		query.Set("locale", "en")
	}
	if !query.Has("vsn") {
		query.Set("vsn", "2.0.0")
	}
	u.RawQuery = query.Encode()
	return u.String(), nil
}

func (c *Config) ParseFromViper(v *viper.Viper) (err error) {
	c.BlockscoutWebSocketUrl, err = fixBlockscoutWebSocketUrl(common.ViperGetOrDefault(v, "websocket.blockscout-websocket-url", "wss://blockscout.dev-02.bas.ankr.com/socket/websocket"))
	if err != nil {
		return err
	}
	c.EthereumWebSocketUrl = common.ViperGetOrDefault(v, "websocket.ethereum-websocket-url", "wss://mainnet.infura.io/ws/v3/261e53a44e3a49eba19eaa54e4907912")
	return nil
}
