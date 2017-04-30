package main

import (
	log "github.com/Sirupsen/logrus"
)

func handleLog(l string) error {
	log.Info("get log: ", l)
	return nil
}
