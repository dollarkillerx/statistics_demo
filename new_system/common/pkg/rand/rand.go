package rand

import (
	"math/rand"
	"strconv"
	"time"
)

// RangeRand 范围随机数
func RangeRand(min, max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := rand.Intn(max-min) + min
	return randNum
}

// RandNumCode 随机数字码
func RandNumCode(codeLen int) string {
	nums := ""
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < codeLen; i++ {
		t := rand.Intn(9)
		nums += strconv.Itoa(t)
	}
	return nums
}
