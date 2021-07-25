package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"net"
)

func sendUDP() {
	p := make([]byte, 2048)
	conn, err := net.Dial("udp", "192.198.5.25:5500")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	s := "EF1299780000000200"
	data, err := hex.DecodeString(s)
	conn.Write(data)
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%X\n", bytes.Trim(p, "\x00"))
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
