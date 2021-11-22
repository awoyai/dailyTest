package daily_practice

import (
	"fmt"
	"testing"
)

func Test_RandomGroup(t *testing.T) {
	resMap := RandomGroup([]int{1, 2, 3, 4, 5, 6, 7, 8})
	for k, v := range resMap {
		fmt.Printf("%v and %v" ,nameMap[k], nameMap[v])
		fmt.Println()
	}
}