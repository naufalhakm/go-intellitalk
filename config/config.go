package config

import "github.com/spf13/viper"

type Config struct {
	MgoHost     string `mapstructure:"MONGO_HOST"`
	MgoPassword string `mapstructure:"MONGO_PASSWORD"`
	MgoDatabase string `mapstructure:"MONGO_DATABASE"`
}

var ENV *Config

func LoadConfig() (*Config, error) {
	fang := viper.New()

	fang.AddConfigPath(".")
	fang.SetConfigName(".")
	fang.SetConfigType("env")

	fang.AutomaticEnv()

	err := fang.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = fang.Unmarshal(&ENV)
	if err != nil {
		// panic(err)
		return nil, err
	}

	return ENV, nil
}
