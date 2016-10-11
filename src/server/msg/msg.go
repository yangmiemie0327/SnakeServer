package msg

import (
	"github.com/name5566/leaf/network/json"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = json.NewProcessor()

func init() {
	// 这里我们注册了一个 JSON 消息 Login
	Processor.Register(&Login{})
	Processor.Register(&Move{})
}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Login  登录
type Login struct {
	AccountId string
	ThemeType int
	RoomId    int
	RoomW     float32
	RoomH     float32
	StartX    float32
	StartY    float32
}

type Move struct {
	AccountId string
	PosX      float32
	PosY      float32
}
