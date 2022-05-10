package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config stores all the configuration of application.
// The value are read by viper from config file or env variable.
type Config struct {
	Post         string `mapstructure:"POST"`
	Environment  string `mapstructure:"GO_ENV"`
	DbSource     string `mapstructure:"DB_URL"`
	jwtSecretKey string `mapstructure:"JWT_SECRET"`
}

// LoadConfig reads the configuration from config file or env variable
func LoadConfig(path string) *Config {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		logrus.Fatalf("Fatal error config file")
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		logrus.Fatalf("Fatal error unmarshalling config file")
	}

	return &config

}
