package system

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yabon-exe/yoggyebiten/game/model"
)

type Random struct {
	seed int64
}

func NewRandom() *Random {
	return &Random{
		seed: time.Now().UnixNano(),
	}
}

func (random *Random) GetRandFromRect(rect model.Rect) (float64, float64) {
	seed := time.Now().UnixNano()
	seed2 := int64(rand.Intn(1000))
	source := rand.NewSource(seed + seed2)
	r := rand.New(source)
	fx := r.Float64()
	fy := r.Float64()

	fmt.Println(fx)
	fmt.Println(fy)
	fmt.Println("--")

	w, h := rect.GetHW()
	return w*fx + rect.Left, h*fy + rect.Top
}
