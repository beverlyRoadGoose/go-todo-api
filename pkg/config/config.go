package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

var Conf *Config

func init() {
	conf, err := ioutil.ReadFile(os.Getenv("CONFIG_FILE"))
	if err != nil {
		log.Error(`unable to read config file: ` + err.Error())
		return
	}

	err = yaml.Unmarshal(conf, &Conf)
	if err != nil {
		log.Error(`error while unmarshalling config file: ` + err.Error())
	}
}
