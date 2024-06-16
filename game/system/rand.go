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

func (random *Random) GetFloat() float64 {
	return random.getRand().Float64()
}

func (random *Random) GetRangeInt(max int) int {
	return random.getRand().Int() % max
}

func (random *Random) GetRandFromRect(rect model.Rect[float64]) (float64, float64) {

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

	vs := make([]T, len(values))
	allVs := make([]T, len(values))
	copy(vs, values)
	copy(allVs, values)

	return &LimitedRandom[T]{
		Random:    NewRandom(),
		allValues: allVs,
		values:    vs,
	}
}

func (lr *LimitedRandom[T]) SetAllValues(values []T) {
	allVs := make([]T, len(values))
	copy(allVs, values)
	lr.allValues = allVs
	lr.Reset()
}

func (lr *LimitedRandom[T]) Reset() {
	vs := make([]T, len(lr.allValues))
	copy(vs, lr.allValues)
	lr.values = vs
}

func (lr *LimitedRandom[T]) PopRandValue() T {
	r := lr.getRand()
	length := len(lr.values)
	index := r.Int() % length
	value := lr.values[index]
	lr.values = append(lr.values[:index], lr.values[index+1:]...)
	return value
}
