package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/catpie/musdk-go"
	"strconv"
	"strings"
)

func handleLog(l string) error {
	log.Info("get log: ", l)
	arr := strings.Split(l, "\n")
	log.Infof("len: %d ", len(arr))
	// var logs []musdk.UserTrafficLog
	for _, v := range arr {
		s := strings.Split(v, " ")
		log.Debugf("processing  %s  len: %d", v, len(s))
		if len(s) < 8 {
			continue
		}
		var sArr []string
		for _, value := range s {
			//log.Debugf("key: %d value %s",key,value)
			if len(value) != 0 {
				sArr = append(sArr, value)
			}
		}
		log.Debugf("user %s  traffic %s", sArr[7], sArr[4])
		u, err := strconv.Atoi(sArr[7])
		if err != nil {
			log.Error("error on get user id ", err)
			continue
		}
		d, err := strconv.Atoi(sArr[4])
		if err != nil {
			log.Error("error on get traffic ", err)
		}
		l := musdk.UserTrafficLog{
			UserId: int64(u),
			D:      int64(d),
		}
		//logs = append(logs, l)
		Queue.Append(l)
	}
	//log.Info("start update traffic to api", logs)
	//err := client.UpdateTraffic(logs)
	//if err != nil {
	//	log.Errorf("error on update traffic  %s", err.Error())
	//}
	return nil
}
