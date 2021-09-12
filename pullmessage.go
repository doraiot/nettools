package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/fananchong/cstruct-go"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var (
	raAddr = getEnv("raAddr", "255.255.255.255:5500")
)

func pullMessage(pc net.PacketConn) {
	udpAddr, err := net.ResolveUDPAddr("udp4", raAddr)
	checkError(err, "net.ResolveUDPAddr")

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
		s := "EF1299780000000200"
		data, _ := hex.DecodeString(s)
		sendUDP(pc, udpAddr, data)
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

func receiveUDP(addr string, data []byte) {
	// s := "050305664A017109000900000000006400000000006400000000006464094D0881010000000064000100204E2F000000000000004100E2FF4100E2FFF40198087206340800000000180A393031B8A8D4CB353030C3D7BFD8D6C6C6F7B5E7D4B400000000000000000000B92E"
	//05 03 05 66 43 01 8B 09 00 05 00 00 00 00 00 64 00 00 00 00 00 64 00 00 00 00 00 64 8F 09 78 08 5A 01 00 00 00 00 64 00 01 00 20 4E 2E 00 00 00 00 00 00 00 41 00 E2 FF 41 00 E2 FF F4 01 98 08 72 06 34 08 00 00 00 00 00 08 39 30 31 20 31 33 30 30 C3 D7 BF D8 D6 C6 C6 F7 B5 E7 D4 B4 00 00 00 00 00 00 00 00 00 00 00 00 D7 70
	// data, _ := hex.DecodeString(s)
	dataLength := len(data)
	if data[0] == AddBoard { //首字节正确，数据长度计算正确
		crcResult := doCrc16(data[:dataLength-2])
		if crcResult[0] == data[dataLength-1] && crcResult[1] == data[dataLength-2] { //CRC判断
			readData := data[4 : data[3]+4]
			fmt.Printf("%x\n", readData)
			switch data[1] {
			case READ_485BYTE:
				// todo read_byte_deal
				break
			case WRITE_485BYTE:
				break
			case READ_485UINT32:
				break
			case WRITE_485UINT32:
				break
			case READ_struct:
				// read struct deal
				switch data[2] {
				case READ_IBDATA:
					mData := &modDataX{}
					cstruct.Unmarshal(readData, mData)
					// jsonData, _ := json.Marshal(mData)
					b, _ := simplifiedchinese.GBK.NewDecoder().Bytes(mData.name[:])
					bdata := &bData{
						IpAddr:    addr,
						MacAddr:   fmt.Sprintf("%v.%v.%v.%v.%v.%v", mData.mac[0], mData.mac[1], mData.mac[2], mData.mac[3], mData.mac[4], mData.mac[5]),
						Name:      string(bytes.Trim(b, "\x00")),
						RawData:   mData,
						CreatedOn: time.Now().Local().Format("20060102150405"),
					}
					go func(bdata *bData) {
						saveToES(bdata)
					}(bdata)

					// fmt.Printf("名称: [%s]\n", b)
					// fmt.Printf("mac: [%v.%v.%v.%v.%v.%v]\n", mData.mac[0], mData.mac[1], mData.mac[2], mData.mac[3], mData.mac[4], mData.mac[5])
					// fmt.Printf("本安输出电压、电流、保护状态: %vV-%vA-%v\n", float64(mData.ib_ch[0].ib_v)/100, float64(mData.ib_ch[0].ib_i/1000), mData.ib_ch[0].protect_flag)
					// fmt.Printf("电池电压、电流：%vV-%.2fA-%v%%-%v\n", float64(mData.v_bat)/100, float64(mData.i_bat)/1000, mData.bat_100, mData.bat_flag)
					// fmt.Printf("%vV-%vA-%v%%-%v\n", float64(mData.v_bat)/100, float64(mData.iout_bat)/1000, mData.bat_100, mData.bat_flag)
					// fmt.Printf("电池温度/电路板温度： %v℃/%v℃\n", mData.ntc_bat, mData.ntc_board)
				default:
					break
				}

			case WRITE_struct:
				break
			default:
				break
			}
		}
	}
}
