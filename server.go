package main

import (
	log "github.com/Sirupsen/logrus"
	"strings"
)

func handleLog(l string) error {
	log.Info("get log: ", l)
	arr := strings.Split(l, "\n")
	log.Infof("len: %d ", len(arr))
	for _, v := range arr {
		s := strings.Split(v, " ")
		log.Debugf("processing  %s  len: %d", v, len(s))
		if len(s) < 8 {
			continue
		}
		//for key,value := range s{
		//	log.Debugf("key: %d value %s",key,value)
		//}
		log.Debugf("user %s  traffic %s", s[10], s[7])

	}
	return nil
}
