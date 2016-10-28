package gamelogic

import (
	"github.com/golang/protobuf/proto"
	"server/common"
	"server/msg/snake"
)

type RoomData struct {
	RoomId     uint32
	RoomW      float32
	RoomH      float32
	FoodNum    uint32
	FoodKeyMax uint32
	PlayerList map[string]int
	FoodList   map[uint32]*FoodData
}

var RoomList map[uint32]*RoomData
var LastFoodAddTime int64
var LastFoodDelTime int64

func init() {
	RoomList = make(map[uint32]*RoomData)
	LastFoodAddTime = 0
	LastFoodDelTime = 0
	initRoom(1, 500, 500, 100)
	initRoom(2, 500, 500, 100)
}

func initRoom(rId uint32, rW float32, rH float32, fN uint32) {
	rData := new(RoomData)
	rData.RoomId = rId
	rData.RoomW = rW
	rData.RoomH = rH
	rData.FoodNum = fN
	rData.FoodKeyMax = 0
	rData.PlayerList = make(map[string]int)
	rData.FoodList = make(map[uint32]*FoodData)
	RoomList[rId] = rData
}

func UpdateRoom() {
	ms := common.GetNowMillisecond()
	for rId, val := range RoomList {
		var msgdata []*snake.MsgPosStruct
		for playerId, _ := range (*val).PlayerList {
			if date, ok := PlayerList[playerId]; ok {
				if date, ok := date.UpdatePlayer(ms); ok {
					msgdata = append(msgdata, &(snake.MsgPosStruct{
						AccountId: proto.String(playerId),
						PosX:      proto.Float32(date.PosX),
						PosY:      proto.Float32(date.PosY),
					}))
				}
			}
		}
		BroadcastRoom(&snake.MsgAddTargetPos{
			PosList: msgdata,
		}, rId)
	}
	if LastFoodAddTime+1000 < ms {
		AddFood()
		LastFoodAddTime = ms
	}
	if LastFoodDelTime+2000 < ms {
		DelFood()
		LastFoodDelTime = ms
	}
}
func AddFood() {
	for rId, val := range RoomList {
		var msgdata []*snake.MsgFoodStruct
		(*val).FoodNum++
		addFood := int((*val).FoodNum) - len((*val).FoodList)
		for i := 0; i < addFood; i++ {
			key := (*val).FoodKeyMax
			fData := new(FoodData)
			fData.InitFoodData((*val).RoomW, (*val).RoomH)
			(*val).FoodList[key] = fData
			msgdata = append(msgdata, &(snake.MsgFoodStruct{
				Id:     proto.Uint32(key),
				PosX:   proto.Float32(fData.Pos.PosX),
				PosY:   proto.Float32(fData.Pos.PosY),
				Radius: proto.Float32(fData.Pos.Radius),
				Score:  proto.Uint32(fData.Score),
			}))
			(*val).FoodKeyMax++
		}
		if len(msgdata) > 0 {
			BroadcastRoom(&snake.MsgAddFood{
				FoodList: msgdata,
			}, rId)
		}
	}
}
func DelFood() {
	for rId, rData := range RoomList {
		var msgdata []uint32
		for fId, fData := range rData.FoodList {
			if fData.IsDel {
				delete(rData.FoodList, fId)
				msgdata = append(msgdata, fId)
			}
		}
		if len(msgdata) > 0 {
			BroadcastRoom(&snake.MsgDelFood{
				FoodList: msgdata,
			}, rId)
		}
	}
}
func AddPlayer(rId uint32, pId string) snake.TErrorType {
	rData, ok := RoomList[rId]
	if !ok {
		return snake.TErrorType_RoomIdIsNull
	}

	if _, ok := PlayerList[pId]; !ok {
		return snake.TErrorType_PlayerIsNo
	}
	if _, ok := (*rData).PlayerList[pId]; !ok {
		PlayerList[pId].InitPlayerData(rId, rData.RoomW, rData.RoomH)
		RoomList[rId].PlayerList[pId] = len(RoomList[rId].PlayerList) + 1
	}
	return snake.TErrorType_Invalid
}
func RemovePlayer(pId string) (bool, uint32) {
	for rId, val := range RoomList {
		for playerId, _ := range val.PlayerList {
			if playerId == pId {
				delete(val.PlayerList, playerId)
				ResetPlayer(pId)
				return true, rId
			}
		}

	}
	return false, 0
}
