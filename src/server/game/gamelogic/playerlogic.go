package gamelogic

import (
	"github.com/name5566/leaf/gate"
	//"github.com/name5566/leaf/log"
	//"github.com/name5566/leaf/util"
)

type PlayerData struct {
	AccountId string
	Connect   gate.Agent
}

var PlayerList map[string]PlayerData

func init() {

	PlayerList = make(map[string]PlayerData)
}

func LoginPlayer(account_id string, a gate.Agent) {
	PlayerList[account_id] = PlayerData{AccountId: account_id, Connect: a}
}

func Broadcast(m interface{}) {
	for _, value := range PlayerList {
		value.Connect.WriteMsg(m)
	}
}
