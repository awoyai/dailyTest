package daily_practice

import (
	"math/rand"
	"testing"
	"time"
)

func Test_Random(t *testing.T) {
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < 5; i++ {
		println(rand.Intn(2))
	}
}
