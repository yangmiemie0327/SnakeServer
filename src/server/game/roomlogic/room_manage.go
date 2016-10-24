package roomlogic

import (
	"fmt"
	"math/rand"
	"server/game/gamelogic"
	"server/msg/snake"
	"time"
)

type RoomData struct {
	RoomId     uint32
	RoomW      float32
	RoomH      float32
	PlayerList map[string]int
}

var RoomList map[uint32]RoomData

func init() {
	RoomList = make(map[uint32]RoomData)
	initRoom(0, 500, 500)
	initRoom(1, 500, 500)
}

func initRoom(rId uint32, rW float32, rH float32) {
	RoomList[rId] = RoomData{
		RoomId:     rId,
		RoomW:      rW,
		RoomH:      rH,
		PlayerList: make(map[string]int),
	}
}

func Update() {
	for _, val := range RoomList {
		for playerId, _ := range val.PlayerList {
			if date, ok := gamelogic.PlayerList[playerId]; ok {
				(&date).Update()
			}
		}
	}
}
func AddPlayer(rId uint32, pId string) snake.TErrorType {
	rData, ok := RoomList[rId]
	if !ok {
		return snake.TErrorType_RoomIdIsNull
	}
	pData, ok := gamelogic.PlayerList[pId]
	if ok {
		return snake.TErrorType_PlayerIsNo
	}
	if _, ok := rData.PlayerList[pId]; ok {
		return snake.TErrorType_PlayerInRoom
	}
	pData.RoomId = rId
	pData.PosList = append(pData.PosList, gamelogic.PosData{
		PosX: rand.Float32() * rData.RoomW * 0.8,
		PosY: rand.Float32() * rData.RoomH * 0.8,
		Time: time.Now().Unix(),
	})
	RoomList[rId].PlayerList[pId] = len(RoomList[rId].PlayerList) + 1
	return snake.TErrorType_Invalid
}
func RemovePlayer(pId string) (bool, uint32) {
	for rId, val := range RoomList {
		for playerId, _ := range val.PlayerList {
			if playerId == pId {
				delete(val.PlayerList, playerId)
				return true, rId
			}
		}

	}
	return false, 0
}
