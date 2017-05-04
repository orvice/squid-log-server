package main

import (
	log "github.com/Sirupsen/logrus"
	"strings"
)

func handleLog(l string) error {
	log.Info("get log: ", l)
	arr := strings.Split(l, "\n")
	log.Infof("len: %d ", len(arr))
	return nil
}
