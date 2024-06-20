package model

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Vector2d struct {
	v *mat.VecDense
}

func NewVector2d[T int | uint | float32 | float64](x T, y T) *Vector2d {
	return &Vector2d{
		v: mat.NewVecDense(2, []float64{float64(x), float64(y)}),
	}
}

func NewDistanceVector2d[T int | uint | float32 | float64](start Vertex[T], end Vertex[T]) *Vector2d {
	return &Vector2d{
		v: mat.NewVecDense(2, []float64{float64(end.X - start.X), float64(end.Y - start.Y)}),
	}
}

func NewVector2dFromDegree(degree float64) *Vector2d {

	// 度数をラジアンに変換
	radian := degree * math.Pi / 180.0
	// X軸とY軸の成分を計算
	x := math.Sin(radian)
	y := math.Cos(radian)

	return NewVector2d[float64](x, y)
}

func (vec *Vector2d) GetX() float64 {
	return vec.v.AtVec(0)
}

func (vec *Vector2d) GetY() float64 {
	return vec.v.AtVec(1)
}

func (vec *Vector2d) Get() (float64, float64) {
	return vec.v.AtVec(0), vec.v.AtVec(1)
}

func (vec *Vector2d) SetX(x float64) {
	vec.v.SetVec(0, x)
}

func (vec *Vector2d) SetY(y float64) {
	vec.v.SetVec(1, y)
}

func (vec *Vector2d) Set(x float64, y float64) {
	vec.SetX(x)
	vec.SetY(y)
}

func (vec *Vector2d) Normalize() {
	norm := vec.v.Norm(2)
	x := vec.GetX() / norm
	y := vec.GetY() / norm
	vec.SetX(x)
	vec.SetY(y)
}
