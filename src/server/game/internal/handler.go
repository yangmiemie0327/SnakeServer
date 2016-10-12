package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/game/gamelogic"
	//"server/msg"
	"server/msg/snake"
)

func init() {
	// 向当前模块（game 模块）注册 Login 消息的消息处理函数 handleLogin
	handler(&snake.Login{}, handleLogin)
	handler(&snake.Move{}, handleMove)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {
	// 收到的 Login 消息
	m := args[0].(*snake.Login)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	gamelogic.LoginPlayer(m.GetAccountId(), a)

	// 给发送者回应一个 Login 消息
	a.WriteMsg(&snake.Login{
		RoomId: proto.Int32(1),
		RoomW:  proto.Float32(0.1),
		RoomH:  proto.Float32(0.2),
		StartX: proto.Float32(0.3),
		StartY: proto.Float32(0.4),
	})
}

func handleMove(args []interface{}) {
	//收到消息
	m := args[0].(*snake.Move)
	// 消息的发送者
	//a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("Move %v %v %v", m.AccountId, m.PosX, m.PosY)

	gamelogic.Broadcast(&snake.Move{
		AccountId: proto.String("abc"),
		PosX:      proto.Float32(0.3),
		PosY:      proto.Float32(0.4),
	})

}
