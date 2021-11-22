package daily_practice

import (
	"math/rand"
	"time"
)

var nameMap = map[int]string {
	1: "刘晓霖",
	2: "曾均宝",
	3: "何亚康",
	4: "黄晓炜",
	5: "麦文建",
	6: "王敬朋",
	7: "吴野",
	8: "杨晓南",
}

func RandomGroup(rands []int) map[int]int {
	length := len(rands)/2
	resMap := make(map[int]int, length)
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < length; i++ {
		index1 := rand.Intn(len(rands))
		rand1 := rands[index1]
		rands = append(rands[:index1], rands[index1+1:]...)
		index2 := rand.Intn(len(rands))
		rand2 := rands[index2]
		rands = append(rands[:index2], rands[index2+1:]...)
		resMap[rand1] = rand2
	}
	return resMap
}