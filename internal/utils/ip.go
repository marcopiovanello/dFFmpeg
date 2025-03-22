package utils

import (
	"log"
	"net"
)

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "0.0.0.0:2080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
