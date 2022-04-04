package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

// Config
type Config struct {
	JWTConfig    JWTConfig    `yaml:"JWTConfig"`
	DBConfig     DBConfig     `yaml:"DBConfig"`
	Logger       Logger       `yaml:"Logger"`
	ServerConfig ServerConfig `yaml:"ServerConfig"`
}

// ServerConfig
type ServerConfig struct {
	AppVersion       string `yaml:"AppVersion"`
	Mode             string `yaml:"Mode"`
	RouterPrefix     string `yaml:"RouterPrefix"`
	Debug            bool   `yaml:"Debug"`
	Port             int    `yaml:"Port"`
	TimeoutSecs      int    `yaml:"TimeoutSecs"`
	ReadTimeoutSecs  int    `yaml:"ReadTimeoutSecs"`
	WriteTimeoutSecs int    `yaml:"WriteTimeoutSecs"`
}

// JWTConfig
type JWTConfig struct {
	SecretKey   string `yaml:"SecretKey"`
	SessionTime int    `yaml:"SessionTime"`
}

// DBConfig
type DBConfig struct {
	DataSourceName string `yaml:"DataSourceName"`
	MaxOpen        int    `yaml:"MaxOpen"`
	MaxIdle        int    `yaml:"MaxIdle"`
	MaxLifeTime    int    `yaml:"MaxLifeTime"`
}

// Logger
type Logger struct {
	Development bool   `yaml:"Development"`
	Encoding    string `yaml:"Encoding"`
	Level       string `yaml:"Level"`
}

func LoadConfig(filename string) (*Config, error){
	v:= viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err:= v.ReadInConfig(); err!=nil{
		if _, ok := err.(viper.ConfigFileNotFoundError); ok{
			return nil, errors.New("config file not found.")
		}
		return nil, err
	}

	var c Config
	err := v.Unmarshal(&c)
	if err!=nil{
		log.Println("unable to decode into struct", err)
		return nil, err
	}

	return &c, nil
}