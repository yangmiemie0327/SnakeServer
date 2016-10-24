package internal

import (
	"fmt"
	"github.com/name5566/leaf/module"
	"server/base"
	"server/game/roomlogic"
	"time"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)
var IsUpdate bool

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	IsUpdate = true
	timeupdate()
}

func (m *Module) OnDestroy() {
	IsUpdate = false
}

func timeupdate() {
	timer := time.NewTimer(100 * time.Millisecond)

	go func() {
		<-timer.C
		if IsUpdate {
			roomlogic.Update()
			timeupdate()
		}
	}()
}
