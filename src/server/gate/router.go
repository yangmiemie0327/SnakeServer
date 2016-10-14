package gate

import (
	"server/game"
	"server/msg"
	"server/msg/snake"
)

func init() {

	// 这里指定消息 Login 路由到 game 模块
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&snake.MsgMsgInit{}, game.ChanRPC)
	msg.Processor.SetRouter(&snake.MsgLogin{}, game.ChanRPC)
	msg.Processor.SetRouter(&snake.MsgMove{}, game.ChanRPC)
	msg.Processor.SetRouter(&snake.MsgRoomInfo{}, game.ChanRPC)
	msg.Processor.SetRouter(&snake.MsgRoomEnter{}, game.ChanRPC)
	msg.Processor.SetRouter(&snake.MsgExitRoom{}, game.ChanRPC)
	msg.Processor.SetRouter(&snake.MsgError{}, game.ChanRPC)
}
