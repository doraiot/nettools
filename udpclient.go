package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	raAddr   = getEnv("raAddr", "255.255.255.255:5500")
	laAddr   = getEnv("laAddr", ":6789")
	internal = getEnv("internal", "60")
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func checkError(err error, funcName string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s-----in func:%s", err.Error(), funcName)
		os.Exit(1)
	}
}
func pullMessage() {
	udpAddr, err := net.ResolveUDPAddr("udp4", raAddr)
	checkError(err, "net.ResolveUDPAddr")

	// localAddr, err := net.ResolveUDPAddr("udp4", laAddr)
	// checkError(err, "net.ResolveUDPAddr")

	// conn, err := net.DialUDP("udp", localAddr, udpAddr)
	// checkError(err, "net.DialUDP")
	//https://github.com/aler9/howto-udp-broadcast-golang
	// defer conn.Close()
	pc, err := net.ListenPacket("udp4", laAddr)
	checkError(err, "net.ListenPacket")

	go func(pc net.PacketConn, udpAddr *net.UDPAddr) {
		keepSendUDP(pc, udpAddr)
	}(pc, udpAddr)

	var buf [512]byte

	for {
		n, addr, err := pc.ReadFrom(buf[:])
		checkError(err, "conn.Read")
		log.Printf("receive from [%v] message length: [%v]\n", addr, n)
		receiveUDP(addr.String(), buf[:])
	}
}

func keepSendUDP(pc net.PacketConn, udpAddr *net.UDPAddr) (bool, error) {
	quit := make(chan struct{})
	internals, _ := strconv.Atoi(internal)
	ticker := time.NewTicker(time.Duration(internals) * time.Second)
	// Keep trying until we're timed out or got a result or got an error
	for {
		sendUDP(pc, udpAddr)
		select {
		// Got a timeout! fail with a timeout error
		case <-quit:
			ticker.Stop()
			return false, errors.New("Stop")
		// Got a tick, we should check on doSomething()
		case <-ticker.C:
			continue
		}
	}
}
