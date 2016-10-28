package gamelogic

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"math/rand"
	"server/common"
	"server/msg/snake"
)

type PlayerData struct {
	AccountId     string
	RoomId        uint32
	DirectionX    float32
	DirectionY    float32
	Speed         float32
	LastTime      int64
	SurplusLength uint32
	PosList       []PosData
	Connect       gate.Agent
}

var PlayerList map[string]*PlayerData

func init() {
	PlayerList = make(map[string]*PlayerData)
}

func LoginPlayer(account_id string, a gate.Agent) {
	if pData, ok := PlayerList[account_id]; ok {
		pData.Connect.WriteMsg(&snake.MsgError{
			ErrorIdx: proto.Uint32(uint32(snake.TErrorType_OtherLogin)),
		})
		PlayerList[account_id].Connect = a
	} else {
		pData := new(PlayerData)
		pData.AccountId = account_id
		pData.Connect = a
		PlayerList[account_id] = pData
	}
}

func Broadcast(m interface{}) {
	for _, value := range PlayerList {
		value.Connect.WriteMsg(m)
	}
}

func BroadcastRoom(m interface{}, rId uint32) {
	for _, value := range PlayerList {
		if value.RoomId == rId {
			value.Connect.WriteMsg(m)
		}
	}
}

func (p *PlayerData) InitPlayerData(rId uint32, rW float32, rH float32) {
	p.RoomId = rId
	p.DirectionX = 1
	p.DirectionY = 0
	p.Speed = 1
	p.SurplusLength = 5
	p.LastTime = common.GetNowMillisecond()
	pos := PosData{}
	pos.InitPosData(rand.Float32()*rW*0.8, rand.Float32()*rH*0.8)
	p.PosList = append(p.PosList, pos)
}

func ResetPlayer(pId string) {
	if pData, ok := PlayerList[pId]; ok {
		pData.RoomId = 0
	}
}

func (p *PlayerData) UpdatePlayer(nowTime int64) (PosData, bool) {
	if len(p.PosList) < 1 {
		return PosData{}, false
	}
	//添加点
	nextX := p.PosList[0].PosX + float32(float64(p.DirectionX)*float64(nowTime-p.LastTime)/1000)
	nextY := p.PosList[0].PosY + float32(float64(p.DirectionY)*float64(nowTime-p.LastTime)/1000)
	tempData := PosData{}
	tempData.InitPosData(nextX, nextY)
	rear := append([]PosData{}, p.PosList[0:]...)
	p.PosList = append(p.PosList[0:0], tempData)
	p.PosList = append(p.PosList, rear...)
	//添加点end

	//过长删除
	if len(p.PosList) > int(p.SurplusLength+3) {
		p.PosList = p.PosList[:p.SurplusLength+3]
	}
	p.LastTime = nowTime
	return tempData, true
}
