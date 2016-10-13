package internal

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	//"github.com/name5566/leaf/log"
	"reflect"
	"server/game/gamelogic"
	"server/msg"
	"server/msg/snake"
)

var MsgList []*snake.MsgMsgData

func init() {
	// 向当前模块（game 模块）注册 Login 消息的消息处理函数 handleLogin
	handler(&snake.MsgMsgInit{}, handleMsgInit)
	handler(&snake.MsgLogin{}, handleLogin)
	handler(&snake.MsgMove{}, handleMove)
	handler(&snake.MsgRoomInfo{}, handleRoomInfo)
	handler(&snake.MsgExitRoom{}, handleExitRoom)
	handler(&snake.MsgError{}, handleError)

	msg.Processor.Range(func(id uint16, t reflect.Type) {
		i := new(snake.MsgMsgData)
		i.MsgId = proto.Uint32(uint32(id))
		i.MsgName = proto.String(t.String())
		MsgList = append(MsgList, i)
	})
	fmt.Print("\n", MsgList, "\n")
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleMsgInit(args []interface{}) {
	//m := args[0].(*snake.MsgMsgInit)
	//a := args[1].(gate.Agent)

	//msgData := snake.MsgMsgInit
}

func handleLogin(args []interface{}) {
	m := args[0].(*snake.MsgLogin)
	a := args[1].(gate.Agent)

	gamelogic.LoginPlayer(m.GetAccountId(), a)
	a.WriteMsg(&snake.MsgLogin{})
}

func handleMove(args []interface{}) {
	//m := args[0].(*snake.MsgMove)

	gamelogic.Broadcast(&snake.MsgMove{})
}

func handleRoomInfo(args []interface{}) {
}

func handleExitRoom(args []interface{}) {
}
func handleError(args []interface{}) {
}
