package utils

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the appilcation.
// The values are read by viper from a config file or enviroment variables.
type Config struct {
	DBDriver                   string        `mapstructure:"DB_DRIVER"`
	DBSource                   string        `mapstructure:"DB_SOURCE"`
	ServerAddress              string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey          string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	ThirdwebPrivateKey         string        `mapstructure:"THIRDWEB_PRIVATE_KEY"`
	MarketplaceContractAddress string        `mapstructure:"MARKETPLACE_CONTRACT_ADDRESS"`
	MetaMaskAddress            string        `mapstructure:"METAMASK_ADDRESS"`
	AccessTokenDuration        time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration       time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig reads configuration from config file or enviroment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	fmt.Println(config)
	return
}
