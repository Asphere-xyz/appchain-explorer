package staking

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

type Config struct {
	Eth1Url         string
	StakingContract common.Address
}

func (c *Config) ParseFromViper(v *viper.Viper) error {
	c.Eth1Url = viperGetOrDefault(v, "staking.eth1-url", "https://rpc.dev-01.bas.ankr.com/")
	rawStakingContract := viperGetOrDefault(v, "staking.staking-contract", "0x0000000000000000000000000000000000001000")
	stakingContract := common.HexToAddress(rawStakingContract)
	if stakingContract == (common.Address{}) {
		return fmt.Errorf("bad staking contract address (%s)", rawStakingContract)
	}
	c.StakingContract = stakingContract
	return nil
}

func viperGetOrDefault(v *viper.Viper, key string, defaultValue string) string {
	if res := v.Get(key); res == nil {
		return defaultValue
	}
	return v.GetString(key)
}
