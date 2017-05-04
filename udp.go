package main

import (
	log "github.com/Sirupsen/logrus"
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
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	log.Debugf("udp conn from %s  buf: %s ", addr, string(buf[:]))
}
