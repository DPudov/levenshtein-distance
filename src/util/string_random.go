package util

import (
	"math"
	"math/rand"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func RandomArray(len int) [] int {
	res := make([]int, len)

	for i := 0; i < len; i++ {
		res[i] = rand.Intn(math.MaxInt64)
	}
	return res
}

func Max(arr []int) (int, int) {
	count := 0
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			count++
		}
	}
	return count, max
}
