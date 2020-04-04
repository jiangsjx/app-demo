package kit

import "github.com/spf13/viper"

const DefaultConfig = "config.yml"

type Config struct {
	App struct {
		Port string
	}

	Log struct {
		Debug bool
		Path  string
	}
}

func NewConfig(file string) (*Config, error) {
	var config Config

	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
