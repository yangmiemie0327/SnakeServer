package gate

import (
	"server/game"
	"server/msg"
	"server/msg/snake"
)

func init() {

	// 这里指定消息 Login 路由到 game 模块
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&snake.Login{}, game.ChanRPC)
	msg.Processor.SetRouter(&snake.Move{}, game.ChanRPC)
}
