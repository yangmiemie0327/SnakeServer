package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
)

func init() {
	// 向当前模块（game 模块）注册 Login 消息的消息处理函数 handleLogin
	handler(&msg.Login{}, handleLogin)
	handler(&msg.Move{}, handleMove)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {
	// 收到的 Login 消息
	m := args[0].(*msg.Login)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("Login %v %v", m.AccountId, m.ThemeType)

	// 给发送者回应一个 Login 消息
	a.WriteMsg(&msg.Login{
		RoomId: 1,
		RoomW:  0.1,
		RoomH:  0.2,
		StartX: 0.3,
		StartY: 0.4,
	})
}

func handleMove(args []interface{}) {
	//收到消息
	m := args[0].(*msg.Move)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("Move %v %v %v", m.AccountId, m.PosX, m.PosY)

	// 给发送者回应一个 Move 消息
	a.WriteMsg(&msg.Move{
		AccountId: "abc",
		PosX:      0.3,
		PosY:      0.4,
	})
}
