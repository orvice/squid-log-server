package main

import (
	log "github.com/Sirupsen/logrus"
	"strings"
)

func handleLog(l string) error {
	log.Info("get log: ", l)
	arr := strings.Split(l, "\n")
	log.Infof("len: %d ", len(arr))
	for _,v := range arr {
		s := strings.Split(v,",")
		if len(s) < 10{
			continue
		}
		log.Debugf("user %s  traffic %s",s[7],s[4])
	}
	return nil
}
