package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type APPConfig struct {
	ServerUrl string
	FrontUrl  string
}

type DBConfig struct {
	User     string
	Host     string
	Database string
	Password string
	Port     string
}

type Config struct {
	ENV string `yaml:"ENV"`
	APP struct {
		ServerUrlDev  string `yaml:"SERVER_URL_DEV"`
		ServerUrlProd string `yaml:"SERVER_URL_PROD"`
		FrontUrlDev   string `yaml:"FRONT_URL_DEV"`
		FrontUrlProd  string `yaml:"FRONT_URL_PROD"`
	} `yaml:"APP"`
	DB struct {
		USER     string `yaml:"USER"`
		HOST     string `yaml:"HOST"`
		DATABASE string `yaml:"DATABASE"`
		PASSWORD string `yaml:"PASSWORD"`
		PORT     string `yaml:"PORT"`
	} `yaml:"DB"`
}

func (config Config) getAppConfig() *APPConfig {
	if config.ENV == "DEV" {
		return &APPConfig{
			config.APP.ServerUrlDev,
			config.APP.FrontUrlDev,
		}
	}
	return &APPConfig{
		config.APP.ServerUrlProd,
		config.APP.FrontUrlProd,
	}
}
func (config Config) getDBConfig() *DBConfig {
	return &DBConfig{
		config.DB.USER,
		config.DB.HOST,
		config.DB.DATABASE,
		config.DB.PASSWORD,
		config.DB.PORT,
	}
}

var config *Config
var AppConfig *APPConfig
var DbConfig *DBConfig

func ReadConfig() error {
	f, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			fmt.Println("Error closing config:\n", err)
		}
	}(f)

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	config = &cfg
	AppConfig = config.getAppConfig()
	DbConfig = config.getDBConfig()
	return nil
}
