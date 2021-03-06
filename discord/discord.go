package discord

import (
	"net"
)

func findAddress() (addr string, err error) {
	var udpAddr *net.UDPAddr
	if udpAddr, err = net.ResolveUDPAddr("udp", "www.internic.net:80"); err != nil {
		return
	}
	var udpConn *net.UDPConn
	if udpConn, err = net.DialUDP("udp", nil, udpAddr); err != nil {
		return
	}
	addr = udpConn.LocalAddr().String()
	return
}
