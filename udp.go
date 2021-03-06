package main

import (
	log "github.com/sirupsen/logrus"
	"net"
)

func startUdpServer() {
	service := Conf.UdpBind
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		log.Panic(err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Panic(err)
	}
	for {
		handleUdpConn(conn)
	}
}

func handleUdpConn(conn *net.UDPConn) {
	var buf [5120]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		log.Errorf("len %d  error: %s", n, err.Error())
		return
	}
	log.Debugf("udp conn from %s  ", addr)
	go handleLog(string(buf[:n]))
}
