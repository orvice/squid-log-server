package main

import (
	log "github.com/cihub/seelog"
)

func handleLog(l string) error {
	log.Info("get log: ", l)
	return nil
}
