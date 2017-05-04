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
	TcpBind  string
	UdpBind  string
	NodeId   string
}

func InitCfg() {
	Conf.Protocol = env("PROTOCOL", "tcp")
	Conf.TcpBind = env("TCP_BIND", ":8801")
	Conf.UdpBind = env("UDP_BIND", ":8802")
	Conf.NodeId = env("NODE_ID","1")
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
