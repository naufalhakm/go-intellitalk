package config

import "github.com/spf13/viper"

type Config struct {
	MgoHost     string `mapstructure:"MONGO_HOST"`
	MgoPassword string `mapstructure:"MONGO_PASSWORD"`
	MgoDatabase string `mapstructure:"MONGO_DATABASE"`
	PortServer  string `mapstructure:"PORT_SERVER"`
}

var ENV *Config

func LoadConfig() {
	fang := viper.New()

	fang.AddConfigPath(".")
	fang.SetConfigName("app")
	fang.SetConfigType("env")

	fang.AutomaticEnv()

	err := fang.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = fang.Unmarshal(&ENV)
	if err != nil {
		panic(err)
	}

}
