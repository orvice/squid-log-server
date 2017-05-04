package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func Init() {
	log.SetOutput(os.Stdout)

	log.Info("boot....")
	InitCfg()
}

func main() {
	Init()
	go startTcpServer()
	go startUdpServer()
}
