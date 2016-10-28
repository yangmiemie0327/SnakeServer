package gamelogic

import (
	"server/msg/snake"
)

type PosData struct {
	PosX   float32
	PosY   float32
	Radius float32
	Shape  snake.TPosShapeType
}

func (p *PosData) InitPosData(px float32, py float32) PosData {
	p.PosX = px
	p.PosY = py
	p.Radius = 1
	p.Shape = snake.TPosShapeType_Round
	return *p
}

func PosOverlapping(a PosData, b PosData) bool {
	return false
}
