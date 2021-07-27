package main

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
	pullMessage()
}
