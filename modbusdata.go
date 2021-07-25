package main

import (
	"encoding/binary"

	"github.com/howeyc/crc16"
)

const (
	AddBoard      = 0x05
	READ_485BYTE  = 0x01
	WRITE_485BYTE = 0x11
	/**********************************READ UINT32***********************************/
	READ_485UINT32 = 0x02
	/**********************************READ UINT32***********************************/

	READ_ADD = 0x01

	READ_CHANNEL1_V = 0x02
	READ_CHANNEL2_V = 0x03
	READ_CHANNEL3_V = 0x04
	READ_CHANNEL4_V = 0x05

	READ_CHANNEL1_I = 0x06
	READ_CHANNEL2_I = 0x07
	READ_CHANNEL3_I = 0x08
	READ_CHANNEL4_I = 0x09

	READ_CHANNEL1_STA = 0x0A
	READ_CHANNEL2_STA = 0x0B
	READ_CHANNEL3_STA = 0x0C
	READ_CHANNEL4_STA = 0x0D

	READ_DC_OR_BAT_OUT = 0x11

	READ_DCIN_V = 0x12

	READ_BAT_I          = 0x13
	READ_BAT_V          = 0x14
	READ_BATIOUT        = 0x15
	READ_BAT_Percentage = 0x16 //百分比
	READ_BAT_FLAG       = 0x17 //电池输入输出拔掉状态

	READ_TEMP_BAT     = 0x21
	READ_TEMP_BOARD   = 0x22
	READ_TEMP_BOARD_F = 0x23
	READ_TEMP_BAT_F   = 0x24

	READ_TEMP_BOARD_L  = 0x25
	READ_TEMP_BOARD_H  = 0x26
	READ_TEMP_SENSER_L = 0x27
	READ_TEMP_SENSER_H = 0x28

	READ_IP   = 0x31 //IP
	READ_DIP  = 0x32 //网关
	READ_U    = 0x33 //掩码
	READ_PORT = 0x34 //端口

	/**********************************WRITE UINT32***********************************/
	WRITE_485UINT32 = 0x12
	/**********************************WRITE UINT32***********************************/
	SET_ONOFF_CH1 = 0x20
	SET_ONOFF_CH2 = 0x21
	SET_ONOFF_CH3 = 0x22
	SET_ONOFF_CH4 = 0x23

	SET_BAT_DC_OUT = 0x24 //市电和电池强制切换

	SET_READ_IP   = 0x32
	SET_READ_DIP  = 0x33
	SET_READ_U    = 0x34
	SET_READ_PORT = 0x35

	SET_TEMP_BOARD_L  = 0x36
	SET_TEMP_BOARD_H  = 0x37
	SET_TEMP_SENSER_L = 0x38
	SET_TEMP_SENSER_H = 0x39

	RESET_IP  = 0x3a
	RESET_REF = 0x3b

	SET_TEMP_SENSER_EN = 0x3C
	SET_BAT_AH         = 0x3D

	CH1_I_REF = 0x40
	CH2_I_REF = 0x41
	CH3_I_REF = 0x42
	CH4_I_REF = 0x43

	CH1_V_REF = 0x44
	CH2_V_REF = 0x45
	CH3_V_REF = 0x46
	CH4_V_REF = 0x47

	BAT_V_REF     = 0x48
	BAT_I_REF     = 0x49
	DCIN_V_REF    = 0x4A
	BATIOUT_V_REF = 0x4B

	SET_ADDR             = 0x4C
	SET_RAT              = 0x4D
	SET_SENT_MODE        = 0x4E //0，被动，1，主动单条，2主动组合
	SET_SENT_TIME        = 0x4F //时间ms
	SET_BAT_OV           = 0x50 //电池过压
	SET_BAT_IN_ONV       = 0x51 //电池插入识别电压
	SET_BAT_GET_BACK_V   = 0x52 //电池充满回头电压
	SET_BAT_OUT_OFF_V    = 0x53 //电池过放保护
	SET_BAT_REON_V       = 0x54 //电池复冲电压
	SET_dc_changge_bat   = 0x55 //电池和市电切换识别电压
	SET_BATIN_COUNT_TING = 0x56 //单次充电最高时长

	READ_AUTO = 0x99 //心跳信号

	/**********************************WRITE sruct***********************************/
	WRITE_struct = 0x13
	/**********************************WRITE sruct***********************************/

	WRITE_NAME       = 0x10
	WRITE_IBDATA_REF = 0x11
	WRITE_IP_BOARD   = 0x12

	/**********************************READ sruct***********************************/
	READ_struct = 0x03
	/**********************************READ sruct***********************************/

	READ_IBDATA = 0x05
	READ_SETING = 0x06

	BAT_OUT = 0x01
	DC_OUT  = 0x00
)

type modDataX struct {
	ib_ch    [4]ibdata
	v_dcin   uint16
	v_bat    uint16
	i_bat    uint16
	iout_bat uint16

	dcORbat_ch uint16

	bat_100          uint16
	bat_flag         uint16
	ntc_bat          int16
	ntc_board        int16
	temp_board_flag  uint16
	temp_senser_falg uint16

	dhcp uint16

	t_bat_h   int16
	t_bat_L   int16
	t_board_h int16
	t_board_l int16

	b_shibie   uint16
	b_guochong uint16
	b_guofang  uint16
	b_fuchong  uint16

	mac  [6]byte
	name [32]byte
}

type ibdata struct {
	ib_i         uint16 //本安输出电流
	ib_v         uint16 //本安输出电压
	protect_flag byte   //保护状态
	change_time  byte   //上次更新时间
}

func doCrc16(data []byte) []byte {
	checksum := ^crc16.ChecksumIBM(data)
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, checksum)
	return b
}
