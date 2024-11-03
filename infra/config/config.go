package config

import "github.com/spf13/viper"

type Config struct {
	App struct {
		Name        string
		Author      string
		Timeout     int
		Development bool
	}
	Bot struct {
		TGToken string
		Debug   bool
	}
	Log struct {
		Level    int
		Encoding string
	}
	PoolSize int
}

var configData *Config

func LoadConfig(fn string) *Config {
	viper.SetConfigFile(fn)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	configData = &Config{
		App: struct {
			Name        string
			Author      string
			Timeout     int
			Development bool
		}{
			Name:        viper.GetString("app.name"),
			Author:      viper.GetString("app.author"),
			Timeout:     viper.GetInt("app.timeout"),
			Development: viper.GetBool("app.development"),
		},
		Bot: struct {
			TGToken string
			Debug   bool
		}{
			TGToken: viper.GetString("bot.tg_token"),
			Debug:   viper.GetBool("bot.debug"),
		},
		Log: struct {
			Level    int
			Encoding string
		}{
			Level:    viper.GetInt("log.level"),
			Encoding: viper.GetString("log.encoding"),
		},
		PoolSize: viper.GetInt("workerpool.size"),
	}

	return configData
}

func GetConfig() *Config {
	return configData
}
