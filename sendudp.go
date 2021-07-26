package main

import (
	"encoding/hex"
	"net"
)

func sendUDP(pc net.PacketConn, udpAddr *net.UDPAddr) {
	s := "EF1299780000000200"
	data, _ := hex.DecodeString(s)
	_, err := pc.WriteTo(data, udpAddr)
	checkError(err, "conn.Write")
}
