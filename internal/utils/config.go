package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var ConfigPath string

type Config struct {
	PostgreSQL    *PostgreSQL    `yaml:"postgresql,omitempty"`
	RestAPIServer *RestAPIServer `yaml:"rest_api_server,omitempty"`
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

type RestAPIServer struct {
	Port string `yaml:"port"`
}

type PostgreSQL struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	User          string `yaml:"user"`
	DBName        string `yaml:"dbname"`
	Password      string `yaml:"password"`
	SSLMode       string `yaml:"sslmode"`
	TimeZone      string `yaml:"timeZone"`
	AutoMigration bool   `yaml:"autoMigration"`
}

func (p *PostgreSQL) PSQLToString() string {
	return "host=" + p.Host + " port=" + p.Port + " dbname=" + p.DBName +
		" user=" + p.User + " password=" + p.Password +
		" sslmode=" + p.SSLMode + " TimeZone=" + p.TimeZone
}
