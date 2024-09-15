package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var Config LocalConfig

func InitConfig() {
	mode := os.Getenv("mode")
	log.Print(mode)
	if mode == "" {
		mode = "local"
	}
	env_file := fmt.Sprintf("config.%v", mode)

	viper.SetConfigName(env_file)
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err := ReadConfig(&Config); err != nil {
		panic(err)
	}
}

func ReadConfigByKey(key string, cfg interface{}) error {
	if err := viper.UnmarshalKey(key, &cfg); err != nil {
		return err
	}
	return nil
}

func ReadConfig(cfg interface{}) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	return nil
}
