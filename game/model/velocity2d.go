package model

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Velocity2d struct {
	v *mat.VecDense
}

func NewVelocity2d(x float64, y float64) *Velocity2d {
	return &Velocity2d{
		v: mat.NewVecDense(2, []float64{x, y}),
	}
}

func NewVelocity2dFromDegree(degree float64) *Velocity2d {

	// 度数をラジアンに変換
	radian := degree * math.Pi / 180.0
	// X軸とY軸の成分を計算
	x := math.Sin(radian)
	y := math.Cos(radian)

	return NewVelocity2d(x, y)
}

func (vec *Velocity2d) GetX() float64 {
	return vec.v.AtVec(0)
}

func (vec *Velocity2d) GetY() float64 {
	return vec.v.AtVec(1)
}

func (vec *Velocity2d) Get() (float64, float64) {
	return vec.v.AtVec(0), vec.v.AtVec(1)
}

func (vec *Velocity2d) SetX(x float64) {
	vec.v.SetVec(0, x)
}

func (vec *Velocity2d) SetY(y float64) {
	vec.v.SetVec(1, y)
}

func (vec *Velocity2d) Scale(scalar float64) {
	vec.v.ScaleVec(scalar, vec.v)
}
