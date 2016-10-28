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
	Score         uint32
	PosList       []PosData
	SnakeNode     []PosData
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
	p.SetScore(0)
	p.LastTime = common.GetNowMillisecond()
	pos := PosData{}
	pos.InitPosData(rand.Float32()*rW*0.8, rand.Float32()*rH*0.8)

	p.PosList = append(p.PosList, pos)
	p.SnakeNode = append(p.SnakeNode, pos)
	p.ResetSnake()
}

func (p *PlayerData) AddScore(s uint32) {
	p.Score += s
	p.SetScore(p.Score)
}

func (p *PlayerData) SetScore(s uint32) {
	p.Score = s
	p.SurplusLength = 5 + p.Score/10
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
	nextX := p.PosList[0].PosX + float32(p.DirectionX*p.Speed)*float32(nowTime-p.LastTime)/1000
	nextY := p.PosList[0].PosY + float32(p.DirectionY*p.Speed)*float32(nowTime-p.LastTime)/1000
	tempData := PosData{}
	tempData.InitPosData(nextX, nextY)
	rear := append([]PosData{}, p.PosList[0:]...)
	p.PosList = append(p.PosList[0:0], tempData)
	p.PosList = append(p.PosList, rear...)
	//添加点end

	p.ResetSnake()

	//过长删除
	if len(p.PosList) > int(p.SurplusLength) {
		p.PosList = p.PosList[:p.SurplusLength]
	}
	p.LastTime = nowTime
	return tempData, true
}

func (p *PlayerData) ResetSnake() {
	lenSnakeNode := len(p.SnakeNode)
	if int(p.SurplusLength) < lenSnakeNode {
		p.SnakeNode = p.SnakeNode[:p.SurplusLength]
	}

	idx := 0
	i := 1
	tempV := p.PosList[idx]
	p.SnakeNode[0].PosX = tempV.PosX
	p.SnakeNode[0].PosY = tempV.PosY
	for ; i < int(p.SurplusLength); i++ {
		if i >= lenSnakeNode {
			pos := PosData{}
			pos.InitPosData(tempV.PosX, tempV.PosY)
			p.SnakeNode = append(p.SnakeNode, pos)
		}

		isHavePos := false
		var dis float32
		dis = 0
		jg := (p.SnakeNode[i-1].Radius + p.SnakeNode[i].Radius) * 0.96
		for ; idx < len(p.PosList); idx++ {
			tempDis := dis
			dis += Distance(tempV, p.PosList[idx])
			if dis > jg {
				tempV = PosLengthToPos(tempV, p.PosList[idx], jg-tempDis)
				isHavePos = true
				break
			} else {
				tempV = p.PosList[idx]
			}
		}
		if !isHavePos {
			tempV.PosX = p.PosList[len(p.PosList)-1].PosX
			tempV.PosY = p.PosList[len(p.PosList)-1].PosY
		}
		p.SnakeNode[i].PosX = tempV.PosX
		p.SnakeNode[i].PosY = tempV.PosY
	}
}
