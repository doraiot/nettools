package main

import (
	"fmt"
	"net"
	"os"
)

var (
	laAddr   = getEnv("laAddr", ":6789")
	internal = getEnv("internal", "600")
)

func main() {
	// s := "EF129978000000"
	// data, err := hex.DecodeString(s)

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("% X\n", data)
	// checksum := ^crc16.ChecksumIBM(data)
	// c := fmt.Sprintf("%04X", checksum)
	// fmt.Println(c[:2])
	// fmt.Println(c[2:])
	// s := "050305664A017109000900000000006400000000006400000000006464094D0881010000000064000100204E2F000000000000004100E2FF4100E2FFF40198087206340800000000180A393031B8A8D4CB353030C3D7BFD8D6C6C6F7B5E7D4B400000000000000000000B92E"
	// data, _ := hex.DecodeString(s)
	// receiveUDP("192.168.5.220:5500", data)
	// pullMessage()

	// localAddr, err := net.ResolveUDPAddr("udp4", laAddr)
	// checkError(err, "net.ResolveUDPAddr")

	// conn, err := net.DialUDP("udp", localAddr, udpAddr)
	// checkError(err, "net.DialUDP")
	//https://github.com/aler9/howto-udp-broadcast-golang
	// defer conn.Close()
	pc, err := net.ListenPacket("udp4", laAddr)
	checkError(err, "net.ListenPacket")
	// pullMessage(pc)

	udpAddr, err := net.ResolveUDPAddr("udp4", "192.168.0.23:5500")
	// udpAddr, err := net.ResolveUDPAddr("udp4", "39.104.185.91:5500")
	checkError(err, "net.ResolveUDPAddr")
	sendMessage(pc, udpAddr, WRITE_NAME, []byte("zhangzs"))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func checkError(err error, funcName string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s-----in func:%s", err.Error(), funcName)
		// os.Exit(1)
	}
}
