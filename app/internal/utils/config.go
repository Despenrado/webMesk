package utils

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

var ConfigPath string

type Config struct {
	PostgreSQL    *PostgreSQLConfig    `yaml:"postgresql,omitempty"`
	RestAPIServer *RestAPIServerConfig `yaml:"rest_api_server,omitempty"`
	RedisConfig   *RedisConfig         `yaml:"redis,omitempty"`
	JWTConfig     *JWTConfig           `yaml:"jwt,omitempty"`
}

func LoadConfig() (*Config, error) {
	configFile, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

type RestAPIServerConfig struct {
	Port string `yaml:"port"`
}

type RedisConfig struct {
	Addr string `yaml:"addr"`
}

type JWTConfig struct {
	TknExpires time.Duration `yaml:"token_expires"`
	JWTKey     string        `yaml:"jwt_key"`
}

type PostgreSQLConfig struct {
	Master struct {
		Host          string `yaml:"host"`
		Port          string `yaml:"port"`
		User          string `yaml:"user"`
		DBName        string `yaml:"dbname"`
		Password      string `yaml:"password"`
		SSLMode       string `yaml:"sslmode"`
		TimeZone      string `yaml:"timeZone"`
		AutoMigration bool   `yaml:"autoMigration"`
	} `yaml:"master"`
	Slave []struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"slave"`
}

// func (p *PostgreSQLConfig) PSQLToString() string {
// 	return "host=" + p.Host + " port=" + p.Port + " dbname=" + p.DBName +
// 		" user=" + p.User + " password=" + p.Password +
// 		" sslmode=" + p.SSLMode + " TimeZone=" + p.TimeZone
// }
