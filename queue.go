package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/catpie/musdk-go"
	"sync"
	"time"
)

var (
	Queue = new(LogQueue)
)

type LogQueue struct {
	logs []musdk.UserTrafficLog
	lock sync.Mutex
}

func (lq *LogQueue) Append(log musdk.UserTrafficLog) error {
	lq.lock.Lock()
	defer lq.lock.Unlock()
	lq.logs = append(lq.logs, log)
	return nil
}

func (lq *LogQueue) SyncToApi() error {
	lq.lock.Lock()

	tmp := make([]musdk.UserTrafficLog, len(lq.logs))
	copy(tmp, lq.logs)
	lq.lock.Unlock()

	log.Infof("proccing tmp logs len %d", len(tmp))

	logMap := map[int64]int64{}
	for _, v := range tmp {
		_, ok := logMap[v.UserId]
		if ok {
			logMap[v.UserId] += v.D + v.U
			continue
		}
		logMap[v.UserId] = v.D + v.U
	}

	tmpLogs := make([]musdk.UserTrafficLog, len(logMap))
	a := 0
	for k, v := range logMap {
		ul := musdk.UserTrafficLog{
			UserId: k,
			D:      v,
		}
		tmpLogs[a] = ul
		a++
	}
	log.Infof("start sync logs to api len: %d", len(tmpLogs))
	return client.UpdateTraffic(tmpLogs)
}

func InitQueue() {

}

func SyncQueue() {
	for {
		err := Queue.SyncToApi()
		if err != nil {
			log.Errorf("error on update traffic  %s", err.Error())
		}
		time.Sleep(time.Second * 60)
	}
}
