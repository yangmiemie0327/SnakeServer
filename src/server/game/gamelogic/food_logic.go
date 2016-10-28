package gamelogic

import (
	"math/rand"
)

type FoodData struct {
	Pos   PosData
	Score uint32
	IsDel bool
}

func (f *FoodData) InitFoodData(rW float32, rH float32) FoodData {
	f.Score = 1
	f.IsDel = false
	f.Pos = PosData{}
	f.Pos.InitPosData(rand.Float32()*rW*0.9, rand.Float32()*rH*0.9)
	return *f
}
