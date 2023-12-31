package configs

import (
	"os"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host      string
	Port      string
	User      string
	Pass      string
	Database  string
	Container string
}

func init() {
	viper.SetDefault("api.port", "8080")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.container", "pgcontainer")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("/app/")
	viper.AddConfigPath("/cmd/server")
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	viper.AddConfigPath(wd)

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:      viper.GetString("database.host"),
		Port:      viper.GetString("database.port"),
		User:      viper.GetString("database.user"),
		Pass:      viper.GetString("database.pass"),
		Database:  viper.GetString("database.name"),
		Container: viper.GetString("database.container"),
	}
	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetPort() string {
	return cfg.API.Port
}
