package array

import (
	"math/rand"
	"time"
)

// Random 随机生成数组
func Random(max, n int) []int {
	rand.Seed(time.Now().Unix())
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, rand.Intn(max))
	}
	return result
}
