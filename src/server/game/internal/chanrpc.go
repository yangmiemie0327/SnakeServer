package internal

import (
	"github.com/name5566/leaf/gate"
	"server/game/roomlogic"
	"time"
)

var IsUpdate bool

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	IsUpdate = true
	timeupdate()
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func timeupdate() {
	skeleton.Go(func() {
		time.Sleep(100 * time.Millisecond)
		roomlogic.Update()
	}, func() {
		timeupdate()
	})
}
