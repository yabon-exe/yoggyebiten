package system

import (
	"math/rand"
	"time"

	"github.com/yabon-exe/yoggyebiten/game/model"
)

type Random struct {
}

func NewRandom() *Random {
	return &Random{}
}

func (random *Random) getRand() *rand.Rand {
	seed := time.Now().UnixNano()
	seed2 := int64(rand.Intn(1000))
	source := rand.NewSource(seed + seed2)
	return rand.New(source)
}

func (random *Random) GetRandFromRect(rect model.Rect) (float64, float64) {

	r := random.getRand()
	fx := r.Float64()
	fy := r.Float64()

	w, h := rect.GetHW()
	return w*fx + rect.Left, h*fy + rect.Top
}

type LimitedRandom[T any] struct {
	*Random
	allValues []T
	values    []T
}

func NewLimitedRandom[T any](values []T) *LimitedRandom[T] {
	return &LimitedRandom[T]{
		Random:    NewRandom(),
		allValues: values,
		values:    values,
	}
}

func (lr *LimitedRandom[T]) SetAllValues(values []T) {
	lr.allValues = values
	lr.Reset()
}

func (lr *LimitedRandom[T]) Reset() {
	lr.values = lr.allValues
}

func (lr *LimitedRandom[T]) PopRandValue() T {
	r := lr.getRand()
	length := len(lr.values)
	index := r.Int() % length
	value := lr.values[index]
	lr.values = append(lr.values[:index], lr.values[index+1:]...)
	return value
}
