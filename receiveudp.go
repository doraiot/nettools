package main

import (
	"fmt"

	"github.com/fananchong/cstruct-go"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func receiveUDP(data []byte) {
	// s := "050305664A017109000900000000006400000000006400000000006464094D0881010000000064000100204E2F000000000000004100E2FF4100E2FFF40198087206340800000000180A393031B8A8D4CB353030C3D7BFD8D6C6C6F7B5E7D4B400000000000000000000B92E"
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
					b, _ := simplifiedchinese.GBK.NewDecoder().Bytes(mData.name[:])
					fmt.Printf("名称: [%s]\n", b)
					fmt.Printf("mac: [%v.%v.%v.%v.%v.%v]\n", mData.mac[0], mData.mac[1], mData.mac[2], mData.mac[3], mData.mac[4], mData.mac[5])
					fmt.Printf("本安输出电压、电流、保护状态: %vV-%vA-%v\n", float64(mData.ib_ch[0].ib_v)/100, float64(mData.ib_ch[0].ib_i/1000), mData.ib_ch[0].protect_flag)
					fmt.Printf("电池电压、电流：%vV-%.2fA-%v%%-%v\n", float64(mData.v_bat)/100, float64(mData.i_bat)/1000, mData.bat_100, mData.bat_flag)
					fmt.Printf("%vV-%vA-%v%%-%v\n", float64(mData.v_bat)/100, float64(mData.iout_bat)/1000, mData.bat_100, mData.bat_flag)
					fmt.Printf("电池温度/电路板温度： %v℃/%v℃\n", mData.ntc_bat, mData.ntc_board)
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
