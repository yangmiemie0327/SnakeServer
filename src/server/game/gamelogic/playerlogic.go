package gamelogic

import (
	"github.com/name5566/leaf/gate"
	//"github.com/name5566/leaf/log"
	//"github.com/name5566/leaf/util"
	"fmt"
	"github.com/golang/protobuf/proto"
	"server/msg/snake"
)

type PosData struct {
	PosX float32
	PosY float32
	Time int64
}
type PlayerData struct {
	AccountId string
	RoomId    uint32
	PosList   []PosData
	Connect   gate.Agent
}

var PlayerList map[string]PlayerData

func init() {
	PlayerList = make(map[string]PlayerData)
}

func LoginPlayer(account_id string, a gate.Agent) {
	if pData, ok := PlayerList[account_id]; ok {
		pData.Connect.WriteMsg(&snake.MsgError{
			ErrorIdx: proto.Uint32(uint32(snake.TErrorType_OtherLogin)),
		})
	}
	PlayerList[account_id] = PlayerData{AccountId: account_id, Connect: a}
}

func Broadcast(m interface{}) {
	for _, value := range PlayerList {
		value.Connect.WriteMsg(m)
	}
}

func (p *PlayerData) Update() {
	fmt.Print(p, "\n")
}
