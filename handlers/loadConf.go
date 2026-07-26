package handlers

import (
	"encoding/json"

	"os"

	"github.com/sirupsen/logrus"
)

type Conf struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

var Cfg Conf

func loadConf() error {
	data, err := os.ReadFile("conf.json")
	Logger("config read", logrus.FatalLevel, err)
	err = json.Unmarshal(data, &Cfg)
	Logger("json problem", logrus.FatalLevel, err)
	return err
}

func init() {
	err := loadConf()
	Logger("config load", logrus.FatalLevel, err)
}
