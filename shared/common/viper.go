package common

import (
	"bytes"
	"encoding/gob"
	"github.com/Ankr-network/ankr-protocol/shared"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var defaultConfigFile = "./config.yaml"

func readFromFile(v *viper.Viper) {
	// default config path
	configPath := v.GetString("config-path")
	if configPath == "" {
		configPath = defaultConfigFile
	}
	// if file doesn't exist just continue
	_, err := os.Stat(configPath)
	if err != nil {
		return
	}
	// set config type from file extension
	ext := filepath.Ext(configPath)[1:]
	v.SetConfigType(ext)
	// set file path
	v.SetConfigFile(configPath)
	// try read config or fatal
	err = v.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read config (%s): %+v", configPath, err)
	}
}

func NewViper() *viper.Viper {
	v := viper.New()
	// enable parse from environment variable
	v.AutomaticEnv()
	// replace "." and "-" with "_" for envs
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	// try read config where its possible
	readFromFile(v)
	return v
}

type ViperConfigProvider struct {
	v *viper.Viper
}

func NewViperConfigProvider(v *viper.Viper) shared.IConfigProvider {
	return &ViperConfigProvider{v: v}
}

func (cp *ViperConfigProvider) Parse(config shared.IConfig) error {
	return config.ParseFromViper(cp.v)
}

func ViperGetOrDefault(v *viper.Viper, key string, defaultValue string) string {
	if v := v.Get(key); v == nil {
		return defaultValue
	}
	return v.GetString(key)
}

func ViperGetOrDefaultUint64(v *viper.Viper, key string, defaultValue uint64) uint64 {
	if v := v.Get(key); v == nil {
		return defaultValue
	}
	return v.GetUint64(key)
}

func ViperGetStringArray(v *viper.Viper, key string, defaultValue []string) []string {
	if v := v.Get(key); v == nil {
		return defaultValue
	}
	return v.GetStringSlice(key)
}

type rawConfigProvider struct {
	src shared.IConfig
}

func (p *rawConfigProvider) Parse(dst shared.IConfig) error {
	buf := bytes.Buffer{}
	err := gob.NewEncoder(&buf).Encode(p.src)
	if err != nil {
		return err
	}
	err = gob.NewDecoder(&buf).Decode(dst)
	if err != nil {
		return err
	}
	return nil
}

func NewRawConfig(config shared.IConfig) shared.IConfigProvider {
	return &rawConfigProvider{src: config}
}
