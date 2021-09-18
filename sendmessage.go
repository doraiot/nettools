package main

import (
	"encoding/hex"
	"fmt"
	"net"
)

func sendMessage(pc net.PacketConn, ipaddr *net.UDPAddr, msgType byte, data []byte) {
	dataLen := len(data)
	bufferData := make([]byte, dataLen+6)
	bufferData[0] = 0x05
	bufferData[1] = WRITE_struct
	bufferData[2] = msgType
	bufferData[3] = byte(dataLen)
	for i := 0; i < dataLen; i++ {
		bufferData[4+i] = data[i]
	}

	crcData := doCrc16(bufferData[0 : dataLen+4])
	bufferData[dataLen+4] = crcData[1]
	bufferData[dataLen+5] = crcData[0]
	fmt.Println(hex.EncodeToString(bufferData))
	sendUDP(pc, ipaddr, bufferData)
}

func sendUDP(pc net.PacketConn, udpAddr *net.UDPAddr, data []byte) {
	_, err := pc.WriteTo(data, udpAddr)
	checkError(err, "conn.Write")
}
