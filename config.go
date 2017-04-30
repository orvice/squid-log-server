package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

var (
	Conf = new(config)
)

type config struct {
	Protocol string
	Bind     string
}

func InitCfg() {
	Conf.Protocol = env("PROTOCOL", "tcp")
	Conf.Bind = env("BIND", ":8801")
}

func env(key, defaultValue string) string {
	v := os.Getenv(key)
	log.Infof("get key: %s, value: %s", key, v)
	if len(v) == 0 {
		log.Info("using default value:", defaultValue)
		return defaultValue
	}
	return v
}
