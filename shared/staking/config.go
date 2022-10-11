package staking

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

type Config struct {
	Eth1Url             string
	StakingContract     common.Address
	ChainConfigContract common.Address
	ChainListUrl        string
	VisibleNetworks     map[uint64]bool
	HiddenNetworks      map[uint64]bool
	BlockTime           uint32
}

func (c *Config) ParseFromViper(v *viper.Viper) (err error) {
	c.Eth1Url = viperGetOrDefault(v, "staking.eth1-url", "ws://rpc.dev-01.bas.ankr.com:8546")
	c.StakingContract, err = viperGetAddressOrDefault(v, "staking.staking-contract", "0x0000000000000000000000000000000000001000")
	if err != nil {
		return err
	}
	c.ChainConfigContract, err = viperGetAddressOrDefault(v, "staking.chain-config-contract", "0x0000000000000000000000000000000000007003")
	if err != nil {
		return err
	}
	c.ChainListUrl = viperGetOrDefault(v, "staking.chain-list-url", "https://raw.githubusercontent.com/node-real/bnbchainlist/main/utils/chains.json")
	hiddenNetworks := viperGetOrDefault(v, "staking.hidden-networks", "56,97")
	c.HiddenNetworks = lo.SliceToMap(strings.Split(hiddenNetworks, ","), func(t string) (uint64, bool) {
		res, _ := strconv.Atoi(t)
		return uint64(res), true
	})
	visibleNetworks := viperGetOrDefault(v, "staking.visible-networks", "")
	c.VisibleNetworks = lo.SliceToMap(strings.Split(visibleNetworks, ","), func(t string) (uint64, bool) {
		res, _ := strconv.Atoi(t)
		return uint64(res), true
	})
	blockTime, err := strconv.Atoi(viperGetOrDefault(v, "staking.block-time", "3000"))
	if err != nil {
		return err
	}
	c.BlockTime = uint32(blockTime)
	return nil
}

func viperGetAddressOrDefault(v *viper.Viper, key, defaultValue string) (common.Address, error) {
	rawValue := viperGetOrDefault(v, key, defaultValue)
	result := common.HexToAddress(rawValue)
	if result == (common.Address{}) {
		return result, fmt.Errorf("bad staking contract address (%s)", rawValue)
	}
	return result, nil
}

func viperGetOrDefault(v *viper.Viper, key, defaultValue string) string {
	if res := v.Get(key); res == nil {
		return defaultValue
	}
	return v.GetString(key)
}
