package gamelogic

import (
	"math"
	"server/msg/snake"
)

type PosData struct {
	PosX   float32
	PosY   float32
	Radius float32
	Width  float32
	Height float32
	Shape  snake.TPosShapeType
}

func (p *PosData) InitPosData(px float32, py float32) PosData {
	p.PosX = px
	p.PosY = py
	p.Radius = 1
	p.Width = 1
	p.Height = 1
	p.Shape = snake.TPosShapeType_Round
	return *p
}

func PosOverlapping(a PosData, b PosData) bool {
	return CheckRound(a, b)
}

func CheckSquare(a PosData, b PosData) bool {
	return false
}

func CheckRound(a PosData, b PosData) bool {
	r2 := a.Radius + b.Radius
	r2 *= r2
	return TwoPosSquare(a, b) < r2
}

//两点的距离
func Distance(a PosData, b PosData) float32 {
	return float32(math.Sqrt(float64(TwoPosSquare(a, b))))
}

//A点向b点l距离的点
func PosLengthToPos(a PosData, b PosData, l float32) PosData {
	dis := Distance(a, b)
	a.PosX = (b.PosX-a.PosX)*l/dis + a.PosX
	a.PosY = (b.PosY-a.PosY)*l/dis + a.PosY
	return a
}

//两点距离的平方
func TwoPosSquare(a PosData, b PosData) float32 {
	x2 := a.PosX - b.PosX
	x2 *= x2
	y2 := a.PosY - b.PosY
	y2 *= y2
	return x2 + y2
}
