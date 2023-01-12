package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type VKConfig struct {
	ClientId     string
	ClientSecret string
}

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
	VK  struct {
		ClientId     string `yaml:"CLIENT_ID"`
		ClientSecret string `yaml:"CLIENT_SECRET"`
	} `yaml:"VK"`
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

func (config Config) getVKConfig() *VKConfig {
	return &VKConfig{
		config.VK.ClientId,
		config.VK.ClientSecret,
	}
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
var VkConfig *VKConfig
var AppConfig *APPConfig
var DbConfig *DBConfig

func ReadConfig() error {
	f, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	config = &cfg
	VkConfig = config.getVKConfig()
	AppConfig = config.getAppConfig()
	DbConfig = config.getDBConfig()
	return nil
}
