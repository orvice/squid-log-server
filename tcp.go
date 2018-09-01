package main

import (
	"fmt"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

func startTcpServer() {

	l, err := net.Listen(Conf.Protocol, Conf.TcpBind)
	if err != nil {
		log.Error("listen error:", err, " net conf: ", Conf)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Error("accept error:", err)
			break
		}
		go handleTcpConn(c)
	}
}

func handleTcpConn(conn net.Conn) error {
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
