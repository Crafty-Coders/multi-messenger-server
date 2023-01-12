package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type VKConfig struct {
	Client_id     string
	Client_secret string
}

type AppConfig struct {
	Server_url string
	Front_url  string
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
		CLIENT_ID     string `yaml:"CLIENT_ID"`
		CLIENT_SECRET string `yaml:"CLIENT_SECRET"`
	} `yaml:"VK"`
	APP struct {
		SERVER_URL_DEV  string `yaml:"SERVER_URL_DEV"`
		SERVER_URL_PROD string `yaml:"SERVER_URL_PROD"`
		FRONT_URL_DEV   string `yaml:"FRONT_URL_DEV"`
		FRONT_URL_PROD  string `yaml:"FRONT_URL_PROD"`
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
		config.VK.CLIENT_ID,
		config.VK.CLIENT_SECRET,
	}
}
func (config Config) getAppConfig() *AppConfig {
	if config.ENV == "DEV" {
		return &AppConfig{
			config.APP.SERVER_URL_DEV,
			config.APP.FRONT_URL_DEV,
		}
	}
	return &AppConfig{
		config.APP.SERVER_URL_PROD,
		config.APP.FRONT_URL_PROD,
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
var VK_config *VKConfig
var APP_config *AppConfig
var DB_config *DBConfig

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
	VK_config = config.getVKConfig()
	APP_config = config.getAppConfig()
	DB_config = config.getDBConfig()
	return nil
}
