package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	//"github.com/name5566/leaf/log"
	"fmt"
	"reflect"
	"server/game/gamelogic"
	"server/game/roomlogic"
	"server/msg"
	"server/msg/snake"
)

var MsgListDate []*snake.MsgMsgData

func init() {
	// 向当前模块（game 模块）注册 Login 消息的消息处理函数 handleLogin
	handler(&snake.MsgMsgInit{}, handleMsgInit)
	handler(&snake.MsgLogin{}, handleLogin)
	handler(&snake.MsgMove{}, handleMove)
	handler(&snake.MsgRoomInfo{}, handleRoomInfo)
	handler(&snake.MsgRoomEnter{}, handleRoomEnter)
	handler(&snake.MsgExitRoom{}, handleExitRoom)
	handler(&snake.MsgError{}, handleError)

	msg.Processor.Range(func(id uint16, t reflect.Type) {
		i := snake.MsgMsgData{
			MsgId:   proto.Uint32(uint32(id)),
			MsgName: proto.String(t.Elem().String()),
		}
		MsgListDate = append(MsgListDate, &i)
	})
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleMsgInit(args []interface{}) {
	fmt.Print("msginit\n")
	a := args[1].(gate.Agent)

	a.WriteMsg(&snake.MsgMsgInit{
		MsgList: MsgListDate,
	})
}

func handleLogin(args []interface{}) {
	m := args[0].(*snake.MsgLogin)
	a := args[1].(gate.Agent)
	fmt.Print("\n^^^^^^^^^^^^^^^\n")
	gamelogic.LoginPlayer(m.GetAccountId(), a)
	a.WriteMsg(&snake.MsgLogin{})
}

func handleMove(args []interface{}) {
	//m := args[0].(*snake.MsgMove)

	gamelogic.Broadcast(&snake.MsgMove{})
}

func handleRoomInfo(args []interface{}) {
	a := args[1].(gate.Agent)
	var data []*snake.MsgRoomData

	for _, v := range roomlogic.RoomList {
		i := snake.MsgRoomData{
			RoomId:      proto.Uint32(v.RoomId),
			RoomW:       proto.Float32(v.RoomW),
			RoomH:       proto.Float32(v.RoomH),
			PlayerCount: proto.Uint32(uint32(len(roomlogic.RoomList))),
		}
		for k, _ := range v.PlayerList {
			i.AccountIdList = append(i.AccountIdList, k)
		}
		data = append(data, &i)
	}

	a.WriteMsg(&snake.MsgRoomInfo{
		RoomList: data,
	})
}

func handleRoomEnter(args []interface{}) {
	m := args[0].(*snake.MsgRoomEnter)
	a := args[1].(gate.Agent)
	err := roomlogic.AddPlayer(m.GetRoomId(), m.GetAccountId())
	if err != snake.TErrorType_Invalid {
		a.WriteMsg(&snake.MsgError{
			ErrorIdx: proto.Uint32(uint32(err)),
		})
		return
	}
	if rData, ok := roomlogic.RoomList[m.GetRoomId()]; ok {
		var data []*snake.MsgPlayerInfo

		for k, _ := range rData.PlayerList {
			if pData, ok := gamelogic.PlayerList[k]; ok {
				msgPdata := snake.MsgPlayerInfo{
					AccountId: proto.String(pData.AccountId),
					RoomId:    proto.Uint32(pData.RoomId),
				}
				for i := 0; i < len(pData.PosList); i++ {
					msgPdata.PosList = append(msgPdata.PosList,
						&snake.MsgPosInfo{
							PosX: proto.Float32(pData.PosList[i].PosX),
							PosY: proto.Float32(pData.PosList[i].PosY),
						})
				}
				data = append(data, &msgPdata)
			}
		}

		a.WriteMsg(&snake.MsgRoomEnter{
			RoomId:      proto.Uint32(rData.RoomId),
			RoomW:       proto.Float32(rData.RoomW),
			RoomH:       proto.Float32(rData.RoomH),
			PlayerCount: proto.Uint32(uint32(len(rData.PlayerList))),
			PlayerList:  data,
		})
	} else {
		a.WriteMsg(&snake.MsgError{
			ErrorIdx: proto.Uint32(uint32(snake.TErrorType_RoomIdIsNull)),
		})
	}
}

func handleExitRoom(args []interface{}) {
	m := args[0].(*snake.MsgExitRoom)
	a := args[1].(gate.Agent)
	if ok, rId := roomlogic.RemovePlayer(m.GetAccountId()); ok {
		gamelogic.BroadcastRoom(&snake.MsgExitRoom{}, rId)
	} else {
		a.WriteMsg(&snake.MsgError{
			ErrorIdx: proto.Uint32(uint32(snake.TErrorType_PlayerNoInRoom)),
		})
	}
}

func handleError(args []interface{}) {
}
