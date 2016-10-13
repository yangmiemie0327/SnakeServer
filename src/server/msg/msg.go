package msg

import (
	"github.com/name5566/leaf/network/protobuf"
	"server/msg/snake"
)

// 使用默认的 proto 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = protobuf.NewProcessor()

func init() {
	// 这里我们注册了一个 JSON 消息 Login
	Processor.Register(&snake.MsgMsgInit{})
	Processor.Register(&snake.MsgLogin{})
	Processor.Register(&snake.MsgMove{})
	Processor.Register(&snake.MsgRoomInfo{})
	Processor.Register(&snake.MsgExitRoom{})
	Processor.Register(&snake.MsgError{})
}
