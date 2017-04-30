package main

import (
	"fmt"
	log "github.com/cihub/seelog"
	"net"
	"time"
)

func main() {
	log.Info("boot....")
	InitCfg()
	l, err := net.Listen(Conf.Protocol, Conf.Bind)
	if err != nil {
		log.Error("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Error("accept error:", err)
			break
		}
		go handleConn(c)
	}
}

func handleConn(conn net.Conn) error {
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	request := make([]byte, 1024)
	defer conn.Close()
	for {
		readLen, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if readLen == 0 {
			break // connection already closed by client
		}
	}
	return handleLog(string(request))
}
