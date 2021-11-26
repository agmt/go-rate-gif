package config

type OpenExchangeConfig struct {
	ApiKey string `mapstructure:"api_key"`
	Base   string `mapstructure:"base"`
}

type GiphyConfig struct {
	ApiKey string `mapstructure:"api_key"`
}

type Config struct {
	Openexchange OpenExchangeConfig `mapstructure:"openexchange"`
	Giphy        GiphyConfig        `mapstructure:"giphy"`
}
