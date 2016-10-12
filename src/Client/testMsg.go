package main

import (
	"encoding/binary"
	"net"
)

func main() {
	sendLogin()
	sendMove()
}
func sendLogin() {

	conn, err := net.Dial("tcp", "192.168.0.180:3563")
	if err != nil {
		panic(err)
	}

	// Login 消息（JSON 格式）
	// 对应游戏服务器 Login 消息结构体
	data := []byte(`{
		"Login": {
			"AccountId": "leaf",
			"ThemeType": 11
		}
	}`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// 发送消息
	conn.Write(m)
}
func sendMove() {

	conn, err := net.Dial("tcp", "192.168.0.180:3563")
	if err != nil {
		panic(err)
	}

	// Login 消息（JSON 格式）
	// 对应游戏服务器 Login 消息结构体
	data := []byte(`{
		"Move": {
			"AccountId": "leaf",
			"PosX": 11,
			"PosY": 10
		}
	}`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// 发送消息
	conn.Write(m)
}
