package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/catpie/musdk-go"
	"os"
	"strconv"
)

var (
	Conf = new(config)
)

type MuCfg struct {
	ApiUri string
	Token  string
	NodeId int
	SType  int
}

type config struct {
	Protocol string
	TcpBind  string
	UdpBind  string
	NodeId   string
	Mu       MuCfg
	// Mu
}

func InitCfg() {
	Conf.Protocol = env("PROTOCOL", "tcp")
	Conf.TcpBind = env("TCP_BIND", ":8801")
	Conf.UdpBind = env("UDP_BIND", ":8802")
	Conf.NodeId = env("NODE_ID", "1")

	Conf.Mu = MuCfg{
		ApiUri: env("API_URI", ""),
		Token:  env("MU_TOKEN", ""),
		NodeId: func() int {
			s := env("NODE_ID", "")
			i, _ := strconv.Atoi(s)
			return i
		}(),
		SType: musdk.TypeHttp,
	}
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
