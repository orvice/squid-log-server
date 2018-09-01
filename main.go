package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.Info("boot....")
	InitCfg()
	InitMu()
}

func waitSignal() {
	var sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP)
	for sig := range sigChan {
		if sig == syscall.SIGHUP {
		} else {
			// is this going to happen?
			log.Infof("caught signal %v, exit", sig)
			os.Exit(0)
		}
	}
}

func main() {
	Init()
	go startTcpServer()
	go startUdpServer()
	go SyncQueue()

	waitSignal()
}
